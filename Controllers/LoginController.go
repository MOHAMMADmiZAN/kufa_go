package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"kufa/DataBase"
	"kufa/KufaSessions"
	"log"
	"net/http"
)

var logAlert = FrontEndAlert{
	Type:         "error",
	ErrorSession: "PasswordError",
	ErrorMessage: "Password Not Valid",
}

func Login(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	//GetAlertMessage(w, r, "passwordErr", "err", "View/Auth/login.gohtml")
	logAlert.GetAlertMessage(w, r, "View/Auth/login.gohtml")
	renderGohtml(w, "Auth/login.gohtml", nil)
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
					logAlert.AddAlertMessage(w, r)
					http.Redirect(w, r, "/login", http.StatusFound)
				} else {
					session, _ := KufaSessions.Store.Get(r, "Go_Session")
					session.Values["authenticated"] = true
					err := session.Save(r, w)
					if err != nil {
						log.Fatalln(err.Error())
					}
					http.Redirect(w, r, "/dashboard", http.StatusFound)
				}

			}

		}
	}

}
func LogOut(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	session, _ := KufaSessions.Store.Get(r, "Go_Session")

	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		log.Fatalln(err.Error())
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
