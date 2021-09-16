package main

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
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
		user := &User{
			Email:    email,
			Password: password,
		}
		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if confirmPassword == password {
			hash := PasswordHash(password)
			insertQuery := "INSERT INTO users(email, password) VALUES ('%s','%s')"
			insertSql := fmt.Sprintf(insertQuery, email, hash)
			insert, err := Db.Query(insertSql)
			if err != nil {
				log.Fatalln(err.Error())
			}
			func(insert *sql.Rows) {
				err := insert.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}(insert)
			fmt.Fprint(w, "Registration Successful")
			//http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

		} else {
			fmt.Fprint(w, "Password Not Match")
		}

	} else {
		fmt.Fprint(w, "notAllaw")
	}

}
