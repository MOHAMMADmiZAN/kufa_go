package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Data struct {
	Name string
	Roll int
}

var tpl *template.Template

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	renderGohtml(w, "dashboard.gohtml", nil)
}
