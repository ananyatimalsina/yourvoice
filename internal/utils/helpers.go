package utils

import (
	"encoding/json"
	"github.com/ananyatimalsina/schema"
	"gorm.io/gorm"
	"reflect"
	"strconv"
)

func RegisterJSONSlicePtr[T any](decoder *schema.Decoder, example []T) {
	decoder.RegisterConverter(example, func(s string) reflect.Value {
		if s == "" || s == "[]" {
			return reflect.ValueOf(example)
		}
		var v []T
		err := json.Unmarshal([]byte(s), &v)
		if err != nil {
			return reflect.Value{}
		}
		return reflect.ValueOf(v)
	})
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
