package Controllers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "index.gohtml", nil)

}
