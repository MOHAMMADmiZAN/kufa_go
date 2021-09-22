package Controllers

import (
	"net/http"
)

func BlogDetails(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "blog-details.gohtml")
}
