package Controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Blog(w http.ResponseWriter, r *http.Request) {
	//parseFiles, err := template.ParseFiles("pages/index.gohtml")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	parseFiles, err := template.ParseFiles("View/blog.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
