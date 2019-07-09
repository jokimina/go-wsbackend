package util

import (
	"reflect"
)

func IndexOf(data interface{}, element interface{}) int {
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		d := reflect.ValueOf(data)
		for i := 0; i < d.Len(); i++ {
			if d.Index(i).Interface() == element {
				return i
			}
		}
	}
	return -1
}

