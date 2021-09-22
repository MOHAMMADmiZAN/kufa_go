package Controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	//renderGohtml(w, "dashboard.gohtml")
	parseFiles, err := template.ParseFiles("View/dashboard.gohtml", "View/master.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		fmt.Println(err.Error())

	}

}
