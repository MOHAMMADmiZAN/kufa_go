package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	renderGohtml(w, "index.gohtml", nil)

}
