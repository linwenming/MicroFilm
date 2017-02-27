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


func NextBillNumber() {
	//http://www.freeyii.com/news/172.html
	//订单号常见的几种方式：
	//1.利用数据库主键值产生一个自增长的订单号（订单号即数据表的主键）
	//2.日期+自增长数字的订单号（比如：2012040110235662）
	//3.产生随机的订单号(65865325365966)
	//4.字母+数字字符串式，字母有包含特别意义，C02356652
	//SELECT auto_increment FROM information_schema.`TABLES` WHERE TABLE_SCHEMA='micro_movie' AND TABLE_NAME='mv_category';
}

