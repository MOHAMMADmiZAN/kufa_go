package Controllers

import (
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	renderGohtml(w, "dashboard.gohtml")

}
