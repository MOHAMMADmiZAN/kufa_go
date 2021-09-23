package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func PortfolioSingle(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	renderGohtml(w, "portfolio-single.gohtml", nil)
}
