package utils

import (
	"reflect"
	"regexp"
	"strings"
)

func CriteriaToMap(obj interface{}) map[string]interface{} {
	// reflect.Value型を取得
	objValue := reflect.ValueOf(obj)
	// ポインタの場合は、ポインタが指すValueを取得
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	// reflect.Type型を取得
	objType := objValue.Type()
	data := make(map[string]interface{})

	isNullList := []string{}
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if strings.Contains(field.Name, "IsNotNull") {
			if objValue.Field(i).Interface() == false {
				isNullList = append(isNullList, ToSnakeCase(strings.Split(field.Name, "IsNotNull")[0]))
			}
		} else {
			fieldValue := objValue.Field(i)
			fieldName := ToSnakeCase(field.Name)
			data[fieldName] = fieldValue.Interface()
		}
	}

	for _, isNullField := range isNullList {
		delete(data, isNullField)
	}

	return data
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
