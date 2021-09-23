package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {

	templates := []string{
		"View/dashboard.gohtml",
		"View/master.gohtml",
	}

	renderMultipleGohtml(w, templates, nil)

}
