package views

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"yourvoice/internal/utils"
	"yourvoice/web/templates"

	"github.com/a-h/templ"
	"gorm.io/gorm"
)

type ModelManagementProps struct {
	Model         any
	SafeModel     any
	PreloadFields []string
	Title         string
	Headers       []string
	MkRow         func(model any) templates.RowProps
}

// TODO: Fix JS errors for modal.Model = null && persistant model data on create
func ModelManagement(w http.ResponseWriter, r *http.Request, db *gorm.DB, props ModelManagementProps) {
	sliceType := reflect.SliceOf(reflect.TypeOf(props.Model))
	modelsSlice := reflect.New(sliceType).Interface()

	// Get search value from query parameter
	//searchValue := r.URL.Query().Get("search")
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

	modelManagerProps := templates.ModelManagerProps{
		Title:   props.Title,
		Headers: props.Headers,
		Rows:    rows,
	}

	modelPage := templates.ModelManager(modelManagerProps)

	if r.Header.Get("AJAX-Target") == "datatable" {
		templ.Handler(modelPage, templ.WithFragments("datatable")).ServeHTTP(w, r)
		return
	}

	if r.Header.Get("AJAX-Target") == "main" {
		templ.Handler(modelPage).ServeHTTP(w, r)
		return
	}

	layout := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return templates.Layout(props.Title).Render(templ.WithChildren(ctx, modelPage), w)
	})
	templ.Handler(layout).ServeHTTP(w, r)
}

func getRows(modelsSlice any, mkRow func(model any) templates.RowProps) []templates.RowProps {
	modelsValue := reflect.ValueOf(modelsSlice).Elem()

	if modelsValue.Kind() != reflect.Slice {
		return []templates.RowProps{}
	}

	length := modelsValue.Len()
	rows := make([]templates.RowProps, length)

	for i := range length {
		item := modelsValue.Index(i).Interface()
		rows[i] = mkRow(item)
	}

	return rows
}
