package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"kufa/DataBase"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	renderGohtml(w, "login.gohtml", nil)
}
func LoginRequest(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
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
