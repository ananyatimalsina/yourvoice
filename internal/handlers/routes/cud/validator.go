package cud

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func ValidateStruct[T any](request T) (string, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		components := make(map[string]string)
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range validationErrors {
				jsonTag := getJSONTag(request, ve.StructField())
				msg := ve.ActualTag()

				components["message-"+jsonTag] = msg
			}
		}

		jsonData, err := json.Marshal(components)
		if err != nil {
			return "", err
		}
		return string(jsonData), err
	}

	return "", nil
}

func getJSONTag(obj any, fieldName string) string {
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
	return fieldName // fallback to struct field name
}
