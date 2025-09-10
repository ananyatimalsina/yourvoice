package cud

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"reflect"
	"strings"
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

func ValidateGorm[T any](request T, err error) (string, error) {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		components := make(map[string]string)
		fields := getFieldsWithUniqueTag(request)

		for _, field := range fields {
			jsonTag := getJSONTag(request, field)
			components["message-"+jsonTag] = err.Error()
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

func getFieldsWithUniqueTag(obj any) []string {
	fields := []string{}
	typ := reflect.TypeOf(obj)
	// If obj is a pointer, get the element type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" && (strings.Contains(gormTag, "unique") || strings.Contains(gormTag, "uniqueIndex") || strings.Contains(gormTag, "unique_constraint")) {
			fields = append(fields, field.Name)
		}
	}
	return fields
}
