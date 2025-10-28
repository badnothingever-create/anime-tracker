package handlers

import (
	"html/template"
	"strings"
)

var funcMap = template.FuncMap{
	"eq": func(a, b interface{}) bool {
		return a == b
	},
	"statusClass": func(status string) string {
		// Пример преобразования: заменить пробелы дефисами и привести к нижнему регистру
		return strings.ToLower(strings.ReplaceAll(status, " ", "-"))
	},
}

var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles(
	"templates/login.html",
	"templates/register.html",
	"templates/index.html",
))
