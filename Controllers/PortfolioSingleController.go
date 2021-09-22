package Controllers

import (
	"net/http"
)

func PortfolioSingle(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "portfolio-single.gohtml")
}
