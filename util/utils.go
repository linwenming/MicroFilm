package util

import (
	"reflect"
	"unicode"
	"bytes"
	"strings"
	"database/sql/driver"
)

func StructMap(value interface{}) map[string]interface{} {

	v := reflect.Indirect(reflect.ValueOf(value))
	m := make(map[string]interface{})

	structValue(m,v)

	return m
}

var (
	typeValuer = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
)

func structValue(m map[string]interface{}, value reflect.Value) {
	if value.Type().Implements(typeValuer) {
		return
	}
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return
		}
		structValue(m, value.Elem())
	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				// unexported
				continue
			}

			tag := field.Tag.Get("db")
			if tag == "-" {
				// ignore
				continue
			}
			if tag == "" {
				// no tag, but we can record the field name
				tag = camelCaseToSnakeCase(field.Name)
			}
			fieldValue := value.Field(i)
			if _, ok := m[tag]; !ok {
				m[tag] = fieldValue.Interface()
			}
			structValue(m, fieldValue)
		}
	}
}

// 提取结构中的db属性
func BuildColumnName(m interface{}) []string{
	v := reflect.TypeOf(m).Elem()

	if v.Kind() == reflect.Struct {
		var value []string
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)

			if field.PkgPath != "" && !field.Anonymous {
				// unexported
				continue
			} else if strings.EqualFold(field.Name,"Id") {
				continue
			}

			tag := field.Tag.Get("db")
			if tag == "-"{
				// ignore
				continue
			}
			if tag == "" {
				// no tag, but we can record the field name
				tag = camelCaseToSnakeCase(field.Name)
			}
			value = append(value,tag)
		}
		return value;
	}

	return []string{}
}

func camelCaseToSnakeCase(name string) string {
	buf := new(bytes.Buffer)

	runes := []rune(name)

	for i := 0; i < len(runes); i++ {
		buf.WriteRune(unicode.ToLower(runes[i]))
		if i != len(runes)-1 && unicode.IsUpper(runes[i+1]) &&
			(unicode.IsLower(runes[i]) || unicode.IsDigit(runes[i]) ||
				(i != len(runes)-2 && unicode.IsLower(runes[i+2]))) {
			buf.WriteRune('_')
		}
	}
	return buf.String()
}




