package helper

import (
	"bytes"
	"html/template"
	"log"
	"reflect"
)

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

func HtmlRenderProcess(templatePath, templateName string, data map[string]interface{}) (*string, error) {
	var tpl *template.Template

	tpl = template.Must(template.ParseFiles(templatePath))

	buf := new(bytes.Buffer)
	err := tpl.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		return nil, err
	}
	html := buf.String()
	return &html, nil
}

func GetJsonTag(fieldName string, obj any) string {
	objType := reflect.TypeOf(obj)
	field, found := objType.FieldByName(fieldName)
	if !found {
		return ""
	}
	return field.Tag.Get("json")
}
