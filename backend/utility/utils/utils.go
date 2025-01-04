package utils

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

func ReflectStruct(ptr interface{}) reflect.Value {
	val := reflect.ValueOf(ptr)
	for val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}
	// fmt.Printf("%v(%v): %v\n", ptr, val.Kind(), val)
	if val.Kind() == reflect.Struct {
		return val
	} else {
		return reflect.ValueOf(nil)
	}
}

func CamelToSnake(str string) string {
	var re = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func SnakeToCamel(str string) string {
	words := strings.Split(str, "_")
	for i := range words {
		if len(words[i]) > 0 {
			words[i] = string(unicode.ToUpper(rune(words[i][0]))) + words[i][1:]
		}
	}
	return strings.Join(words, "")
}
