package helper

import (
	"bytes"
	"html/template"
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

//func HtmlRenderProcess(templatePath string, data map[string]interface{}) (*string, error) {
//	// парсим наш шаблон
//	tmpl := template.Must(template.ParseFiles(templatePath))
//	// создаём пустой буфер для записи
//	buf := new(bytes.Buffer)
//	// рендерим шаблон с данными data в наш буфер buf
//	err := tmpl.Execute(buf, data)
//	if err != nil {
//		return nil, err
//	}
//	html := buf.String()
//	return &html, err
//}
