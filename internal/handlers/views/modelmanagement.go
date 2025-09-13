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
)

type ModelManagementProps struct {
	Model        any
	SafeModel    any
	SearchFields []string
	Title        string
	Headers      []string
	PrepareDB    func(db *gorm.DB) *gorm.DB
	MkRow        func(model any) modelmanagement.RowProps
	ModalProps   modelmanagement.ModalProps
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

	// var managerModel any
	// if props.SafeModel != nil {
	// 	managerModel = props.SafeModel
	// } else {
	// 	managerModel = props.Model
	// }

	if searchQuery != "" {
		for _, field := range props.SearchFields {
			db = db.Or(field+" ILIKE ?", "%"+searchQuery+"%")
		}
	}

	db = db.Offset((size * (page - 1))).Limit(size)

	if err := db.Find(modelsSlice).Error; err != nil {
		http.Error(w, "Failed to retrieve objects", http.StatusInternalServerError)
		return
	}

	rows := getRows(modelsSlice, props.MkRow)

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
