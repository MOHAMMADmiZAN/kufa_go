package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
func Blog(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	parseFiles, err = template.ParseFiles("pages/blog.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func PortfolioSingle(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	parseFiles, err = template.ParseFiles("pages/portfolio-single.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func BlogDetails(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	parseFiles, err = template.ParseFiles("pages/blog-details.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/login.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/login.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	parseFiles, err = template.ParseFiles("pages/register.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func RegisterRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirmPassword")

		_, err := fmt.Fprint(w, email, password, confirmPassword)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		fmt.Fprint(w, "notAllaw")
	}

}
func InboxRequest(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	insertQuery := "INSERT INTO `inbox` (`name`,`email`,`messages`) VALUES ('%s','%s','%s')"
	insertSql := fmt.Sprintf(insertQuery, name, email, message)
	insert, err := Db.Query(insertSql)
	if err != nil {
		log.Fatalln(err)
	}
	func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(insert)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

}
