package utils

import "reflect"

func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	return v.IsZero()
}

func IsNotEmpty(value interface{}) bool {
	if value != nil {
		v := reflect.ValueOf(value)
		return !v.IsZero()
	}
	return false
}
