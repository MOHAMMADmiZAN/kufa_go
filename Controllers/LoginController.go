package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"html/template"
	"kufa/DataBase"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("View/login.gohtml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func LoginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		loginUser := &LoginUser{
			Email:    email,
			Password: password,
		}
		validate := validator.New()
		err := validate.Struct(loginUser)
		if err != nil {
			fmt.Fprint(w, err.(validator.ValidationErrors))
		} else {
			var ReturnUser string
			var ReturnPassword string
			err := DataBase.Db.QueryRow("SELECT email, password FROM users WHERE email=?", email).Scan(&ReturnUser, &ReturnPassword)
			if err == sql.ErrNoRows && err != nil {
				fmt.Println(err)
				http.Redirect(w, r, "/login", http.StatusFound)
			} else {
				err := MatcherHash(ReturnPassword, password)
				if err != nil {
					fmt.Println(err)
					http.Redirect(w, r, "/login", http.StatusFound)
				} else {
					http.Redirect(w, r, "/dashboard", http.StatusFound)
				}

			}

		}
	}

}
