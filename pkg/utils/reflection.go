package utils

import (
	"reflect"
	"strings"
)

func AllFields(t any) []string {
	var fields []string

	st := reflect.TypeOf(t)
	for field := range st.Fields() {
		fields = append(fields, strings.ToLower(field.Name))
	}
	return fields
}
