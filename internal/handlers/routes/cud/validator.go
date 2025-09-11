package cud

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ValidateStruct[T any](request T, ignoreUnique bool) (string, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		components := make(map[string]string)
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range validationErrors {
				// if the gorm tag includes "unique" and we're ignoring unique constraints, skip it
				if ignoreUnique && hasUniqueTag(getGormTag(request, ve.StructField())) {
					continue
				}
				jsonTag := getJSONTag(request, ve.StructField())
				msg := ve.ActualTag()

				components["message-"+jsonTag] = msg
			}
		}

		if len(components) == 0 {
			return "", nil
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

func GetIdsFromAjax(r *http.Request) ([]uint64, error) {
	ids := []uint64{}

	if r.Header.Get("AJAX-Targets") != "" {
		targets := strings.SplitSeq(r.Header.Get("AJAX-Targets"), ",")
		for t := range targets {
			id := strings.TrimPrefix(t, "row-")
			parseUint, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				return nil, err
			}
			ids = append(ids, parseUint)
		}
	}

	if r.Header.Get("AJAX-Target") != "" {
		id := strings.TrimPrefix(r.Header.Get("AJAX-Target"), "row-")
		parseUint, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, parseUint)
	}

	if len(ids) == 0 {
		return nil, errors.New("No IDs found in AJAX headers")
	}

	return ids, nil
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

func getGormTag(obj any, fieldName string) string {
	typ := reflect.TypeOf(obj)
	// If obj is a pointer, get the element type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, found := typ.FieldByName(fieldName)
	if found {
		tag := field.Tag.Get("gorm")
		if tag != "" {
			return tag
		}
	}
	return "" // fallback to empty string if not found
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
		if hasUniqueTag(gormTag) {
			fields = append(fields, field.Name)
		}
	}
	return fields
}

func hasUniqueTag(tag string) bool {
	return strings.Contains(tag, "unique") || strings.Contains(tag, "uniqueIndex") || strings.Contains(tag, "unique_constraint")
}
