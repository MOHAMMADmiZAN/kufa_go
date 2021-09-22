package Controllers

import (
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "blog.gohtml")
}
