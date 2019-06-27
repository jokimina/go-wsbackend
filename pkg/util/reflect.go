package util

import (
	"reflect"
)

type resultItem struct {
	Value reflect.Value
	Field reflect.StructField
	Tag   reflect.StructTag
}

func GetFields(m interface{}, skipStructField bool)(result []resultItem){
	val := reflect.ValueOf(m).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if skipStructField && typeField.Type.Kind() == reflect.Struct{
			continue
		}
		valueField := val.Field(i)
		tag := typeField.Tag
		result = append(result, resultItem{
			Value: valueField,
			Field: typeField,
			Tag: tag,
		})
		//fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
	return
}

func GetFieldByName(e interface{}, field string) reflect.Value {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}
