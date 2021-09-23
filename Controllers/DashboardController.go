package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Data struct {
	Name string
	Roll int
}

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	template := []string{
		"View/master.gohtml",
		"View/dashboard.gohtml",
	}
	renderMulipleGohtml(w, template, nil)
}
