package views

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"yourvoice/internal/utils"
	"yourvoice/web/templates"
	"yourvoice/web/templates/modelmanagement"

	"github.com/a-h/templ"
	"gorm.io/gorm"
	"strings"
)

type ModelManagementProps struct {
	Model         any
	SafeModel     any
	PreloadFields []string
	SearchFields  []string
	Title         string
	Headers       []string
	MkRow         func(model any) modelmanagement.RowProps
	ModalProps    modelmanagement.ModalProps
}

// TODO: Fix JS errors for modal.Model = null && persistant model data on create
func ModelManagement(w http.ResponseWriter, r *http.Request, db *gorm.DB, props ModelManagementProps) {
	ctx := r.Context()
	sliceType := reflect.SliceOf(reflect.TypeOf(props.Model))
	modelsSlice := reflect.New(sliceType).Interface()

	searchQuery := r.URL.Query().Get("search")

	orderBy := r.URL.Query().Get("orderBy")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 {
		size = utils.PageSizeCompact
	}

	totalItems := int64(0)
	db.Model(&props.Model).Count(&totalItems)

	// Fetch parties from database with search filter
	for _, field := range props.PreloadFields {
		db = db.Preload(field)
	}

	if searchQuery != "" {
		for _, field := range props.SearchFields {
			db = db.Or(field+" ILIKE ?", "%"+searchQuery+"%")
		}
	}

	db = getOrder(db, props.Model, orderBy, props.ModalProps.Title, props.Title)

	db = db.Offset((size * (page - 1))).Limit(size)

	if err := db.Find(modelsSlice).Error; err != nil {
		http.Error(w, "Failed to retrieve objects", http.StatusInternalServerError)
		return
	}

	rows := getRows(modelsSlice, props.MkRow)

	// var managerModel any
	// if props.SafeModel != nil {
	// 	managerModel = props.SafeModel
	// } else {
	// 	managerModel = props.Model
	// }

	modelManagerProps := modelmanagement.ModelManagerProps{
		Title:        props.Title,
		Headers:      props.Headers,
		Rows:         rows,
		ModalProps:   props.ModalProps,
		SearchQuery:  searchQuery,
		CurrentOrder: orderBy,
	}

	modelPage := modelmanagement.ModelManager(modelManagerProps)

	if r.Header.Get("AJAX-Target") == "datatable" {
		templ.RenderFragments(ctx, w, modelPage, "datatable")
		return
	}

	if r.Header.Get("AJAX-Target") == "main" {
		modelPage.Render(ctx, w)
		return
	}

	layout := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return templates.Layout(props.Title).Render(templ.WithChildren(ctx, modelPage), w)
	})
	templ.Handler(layout).ServeHTTP(w, r)
}

func getRows(modelsSlice any, mkRow func(model any) modelmanagement.RowProps) []modelmanagement.RowProps {
	modelsValue := reflect.ValueOf(modelsSlice).Elem()

	if modelsValue.Kind() != reflect.Slice {
		return []modelmanagement.RowProps{}
	}

	length := modelsValue.Len()
	rows := make([]modelmanagement.RowProps, length)

	for i := range length {
		item := modelsValue.Index(i).Interface()
		rows[i] = mkRow(item)
	}

	return rows
}

func getOrder(db *gorm.DB, obj any, userField string, singular string, plural string) *gorm.DB {
	desc := ""
	if strings.HasPrefix(userField, "-") {
		desc = " desc"
	}
	userField = strings.ReplaceAll(strings.TrimPrefix(userField, "-"), " ", "")
	singular = strings.ToLower(singular)
	plural = strings.ToLower(plural)

	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, found := typ.FieldByName(userField)
	if !found {
		// Field not found, return db unmodified
		return db
	}

	jsonTag := utils.GetJSONTag(obj, userField)
	if jsonTag == "" {
		return db
	}

	// Check if field is a slice (relation)
	if field.Type.Kind() == reflect.Slice {
		// Example: Candidates []Candidate
		// You might want to support more relations, but here's for Candidates
		// Assumes Candidate has a party_id foreign key
		join := "LEFT JOIN " + jsonTag + " ON " + jsonTag + "." + singular + "_id = " + plural + ".id"
		group := plural + ".id"
		order := "COUNT(" + jsonTag + ".id)" + desc
		return db.Joins(join).Group(group).Order(order)
	} else {
		// Regular field
		return db.Order(jsonTag + desc)
	}
}
