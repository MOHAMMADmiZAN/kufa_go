package Controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {

	templates := []string{
		"View/dashboard.gohtml",
		"View/master.gohtml",
	}
	renderMultipleGohtml(w, templates)
	//fmt.Println(template)
	//files, err := template.ParseFiles("View/dashboard.gohtml", "View/master.gohtml")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//files.Execute(w,nil)
}
