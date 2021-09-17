package Controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("view/index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
