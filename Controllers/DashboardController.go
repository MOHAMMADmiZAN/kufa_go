package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {

	renderWithAuth(w, r, nil, "View/Dashboard/dashboard.gohtml")

}
