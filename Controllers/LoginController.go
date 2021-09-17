package Controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("View/login.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
