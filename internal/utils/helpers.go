package utils

import (
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"reflect"
	"strconv"
)

func GetModelID(model any) uint64 {
	if model == nil {
		return 0
	}

	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return 0
	}

	idField := val.FieldByName("ID")
	if !idField.IsValid() {
		return 0
	}

	switch idField.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return idField.Uint()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(idField.Int())
	case reflect.String:
		u64, err := strconv.ParseUint(idField.String(), 10, 64)
		if err != nil {
			return 0
		}
		return u64
	default:
		return 0
	}
}

func GetJSONTag(obj any, fieldName string) string {
	typ := reflect.TypeOf(obj)
	// If obj is a pointer, get the element type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, found := typ.FieldByName(fieldName)
	if found {
		tag := field.Tag.Get("json")
		if tag != "" {
			return tag
		}
	}
	return ""
}

func BuildRelationshipFieldInputOptions(db *gorm.DB, modelType any) []InputOption {
	sliceType := reflect.SliceOf(reflect.TypeOf(modelType))
	modelsSlice := reflect.New(sliceType).Interface()

	if err := db.Find(modelsSlice).Error; err != nil {
		return []InputOption{}
	}

	modelsValue := reflect.ValueOf(modelsSlice).Elem()

	if modelsValue.Kind() != reflect.Slice {
		return []InputOption{}
	}

	length := modelsValue.Len()
	options := make([]InputOption, length)

	for i := range length {
		item := modelsValue.Index(i)

		// Handle pointer to struct
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		var id uint
		var label string

		// Try to find ID field (ID, id, or first uint field)
		if idField := item.FieldByName("ID"); idField.IsValid() && idField.CanInterface() {
			if idVal, ok := idField.Interface().(uint); ok {
				id = idVal
			}
		} else if idField := item.FieldByName("id"); idField.IsValid() && idField.CanInterface() {
			if idVal, ok := idField.Interface().(uint); ok {
				id = idVal
			}
		} else {
			for j := 0; j < item.NumField(); j++ {
				field := item.Field(j)
				if field.Kind() == reflect.Uint && field.CanInterface() {
					if idVal, ok := field.Interface().(uint); ok {
						id = idVal
						break
					}
				}
			}
		}

		// Try to find label field (Name, Title, Label, or first string field)
		labelFields := []string{"Name", "Title", "Label"}
		for _, fieldName := range labelFields {
			if labelField := item.FieldByName(fieldName); labelField.IsValid() && labelField.CanInterface() {
				if labelVal, ok := labelField.Interface().(string); ok && labelVal != "" {
					label = labelVal
					break
				}
			}
		}

		// Fallback: find first string field if no standard label field found
		if label == "" {
			for j := 0; j < item.NumField(); j++ {
				field := item.Field(j)
				if field.Kind() == reflect.String && field.CanInterface() {
					if strVal := field.Interface().(string); strVal != "" {
						label = strVal
						break
					}
				}
			}
		}

		options[i] = InputOption{
			Value: strconv.Itoa(int(id)),
			Label: label,
		}
	}

	return options
}

func GetJSONString(v any) string {
	jsonStr, err := templ.JSONString(v)
	if err != nil {
		return "null"
	}
	return jsonStr
}
