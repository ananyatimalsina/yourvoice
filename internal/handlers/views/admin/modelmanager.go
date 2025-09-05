package admin

import (
	"context"
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"yourvoice/internal/utils"
	"yourvoice/web/templates/admin"
	"yourvoice/web/templates/admin/components"
	"yourvoice/web/templates/admin/components/basic"
)

type ModelManagerProps struct {
	Model              any
	SafeModel          any
	SearchFields       []string
	PreloadFields      []string
	Headers            []string
	MkRow              func(model any) components.RowProps
	Title              string
	SingularTitle      string
	Icon               string
	RelationshipFields []components.RelationshipField
	Actions            []templ.Component
	Options            [2]bool
}

// TODO: Fix JS errors for modal.Model = null && persistant model data on create
func ModelManager(w http.ResponseWriter, r *http.Request, db *gorm.DB, props ModelManagerProps) {
	ctx := r.Context()

	sliceType := reflect.SliceOf(reflect.TypeOf(props.Model))
	modelsSlice := reflect.New(sliceType).Interface()

	// Get search value from query parameter
	searchValue := r.URL.Query().Get("search")
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

	for _, field := range props.SearchFields {
		db = db.Or(field+" ILIKE ?", "%"+searchValue+"%")
	}

	db = db.Offset((size * (page - 1))).Limit(size)

	if err := db.Find(modelsSlice).Error; err != nil {
		http.Error(w, "Failed to retrieve objects", http.StatusInternalServerError)
		return
	}

	rows := getRows(modelsSlice, props.MkRow)

	var managerModel any
	if props.SafeModel != nil {
		managerModel = props.SafeModel
	} else {
		managerModel = props.Model
	}

	// TODO: handle data row as child if no HX-Request

	modelManagerProps := components.ModelManagerProps{
		Title:         props.Title,
		SingularTitle: props.SingularTitle,
		Icon:          props.Icon,
		SearchValue:   searchValue,
		MainAction: basic.ActionButton(basic.ActionProps{
			Label: "New " + props.SingularTitle,
			Style: utils.ActionStylePrimary,
			Attributes: templ.Attributes{
				"x-on:click": "Object.assign(modal, { Open: true, Model: null})",
			},
		}),
		Headers:    props.Headers,
		Rows:       rows,
		RowActions: props.Actions,
		Modal: components.ModalProps{
			RelationshipFields: props.RelationshipFields,
			Model:              managerModel},
		Pagination: components.PaginationProps{
			CurrentPage:  page,
			TotalItems:   totalItems,
			ItemsPerPage: size,
		},
	}

	dataTableBody := components.TBody(components.TBodyProps{
		Rows:       modelManagerProps.Rows,
		RowActions: modelManagerProps.RowActions,
	}, [3]bool{props.Options[0], props.Options[1], true})

	if r.Header.Get("X-Alpine-Request") != "" {
		// If the request is an Alpine request, render only the DataTable component
		dataTableBody.Render(ctx, w)
		return
	}

	// Create a wrapper component that renders both ModelManager and Modal inside the Layout
	modelPage := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		// Create a component that includes both the ModelManager and the Modal
		pageContent := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return components.ModelManager(modelManagerProps, props.Options).Render(templ.WithChildren(ctx, dataTableBody), w)
		})

		return admin.Layout(props.Title).Render(templ.WithChildren(ctx, pageContent), w)
	})

	// Render the composed component
	templ.Handler(modelPage).ServeHTTP(w, r)
}

func getRows(modelsSlice any, mkRow func(model any) components.RowProps) []components.RowProps {
	modelsValue := reflect.ValueOf(modelsSlice).Elem()

	if modelsValue.Kind() != reflect.Slice {
		return []components.RowProps{}
	}

	length := modelsValue.Len()
	rows := make([]components.RowProps, length)

	for i := range length {
		item := modelsValue.Index(i).Interface()
		rows[i] = mkRow(item)
	}

	return rows
}
