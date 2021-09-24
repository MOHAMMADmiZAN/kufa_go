package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	data := struct {
		Name string
	}{
		Name: "Mizan",
	}
	renderMultipleGohtml(w, data, "View/Dashboard/dashboard.gohtml")

}
