package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	renderGohtml(w, "blog.gohtml", nil)
}
