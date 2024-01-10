package helper

import (
	"log"
	"reflect"
	"task_tracker/internal/domain"
)

type Entities interface {
	domain.User | domain.Position | domain.Department
}

func GetEntityData(data any) ([]string, []any) {
	var fields []string
	var values []any

	a := reflect.ValueOf(data)
	index := reflect.ValueOf(data).Elem().NumField()
	if a.Kind() != reflect.Ptr {
		log.Fatal("wrong type struct")
	}
	for x := 0; x < index; x++ {
		fields = append(fields, reflect.TypeOf(data).Elem().Field(x).Tag.Get("db"))
		values = append(values, reflect.ValueOf(data).Elem().Field(x).Interface())
	}
	return fields, values
}
