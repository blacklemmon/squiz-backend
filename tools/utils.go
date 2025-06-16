package tools

import (
	"reflect"
)

func ArrayFind(slice interface{}, element interface{}) int {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
			panic("ArrayFind : O primeiro argumento deve ser um slice")
	}

	for i := 0; i < v.Len(); i++ {
			if reflect.DeepEqual(v.Index(i).Interface(), element) {
					return i
			}
	}
	return -1
}

/**
/* Do not be mistaken: StrFind evaluates whether a string is present in a slice of strings
*/
func StrFind(value string, values ...string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}