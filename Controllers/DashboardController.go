package Controllers

import (
	"net/http"
)

type Data struct {
	Name string
	Roll int
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "master.gohtml", nil)

	data := Data{Name: "Mizan", Roll: 15}
	renderGohtml(w, "dashboard.gohtml", data)

}
