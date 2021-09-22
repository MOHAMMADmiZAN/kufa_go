package Controllers

import (
	"html/template"
	"net/http"
)

type Data struct {
	Name string
	Roll int
}

var tpl *template.Template

func Dashboard(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "dashboard.gohtml", nil)
}
