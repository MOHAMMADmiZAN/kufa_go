package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func BlogDetails(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	renderGohtml(w, "blog-details.gohtml", nil)
}
