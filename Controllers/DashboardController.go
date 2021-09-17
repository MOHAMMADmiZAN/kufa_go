package Controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	parseFiles, err := template.ParseFiles("View/dashboard.gohtml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
