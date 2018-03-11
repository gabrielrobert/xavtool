package main

import (
	"reflect"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func containsStructFieldValue(slice interface{}, fieldName string, fieldValueToCheck interface{}) bool {

	rangeOnMe := reflect.ValueOf(slice)

	for i := 0; i < rangeOnMe.Len(); i++ {
		s := rangeOnMe.Index(i)
		f := s.FieldByName(fieldName)
		if f.IsValid() {
			fieldAsString := f.Interface().(string)
			fieldValueAsString := fieldValueToCheck.(string)
			if strings.Contains(fieldAsString, fieldValueAsString) {
				return true
			}
		}
	}

	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
