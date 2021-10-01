package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"kufa/DataBase"
	"log"
	"net/http"
)

var registerAlert = FrontEndAlert{
	Type:         "register",
	ErrorMessage: "You Register Successfully",
	ErrorSession: "registerGreen",
}

func Register(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	registerAlert.GetAlertMessage(w, r, "View/Auth/login.gohtml")
	renderGohtml(w, "Auth/register.gohtml", nil)
}
func RegisterRequest(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err.Error())
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("confirmPassword")
		user := &User{
			Email:           email,
			Password:        password,
			ConfirmPassword: confirmPassword,
		}
		validate := validator.New()
		err = validate.Struct(user)
		if err != nil {
			fmt.Fprint(w, err.(validator.ValidationErrors))
		} else {
			hash := PasswordHash(password)
			insertQuery := "INSERT INTO users(email, password) VALUES ('%s','%s')"
			insertSql := fmt.Sprintf(insertQuery, email, hash)
			insert, err := DataBase.Db.Query(insertSql)
			if err != nil {
				log.Fatalln(err)
			} else {
				registerAlert.AddAlertMessage(w, r)

			}
			func(insert *sql.Rows) {
				err := insert.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}(insert)
			registerAlert.GetAlertMessage(w, r, "View/Auth/login.gohtml")
		}

		//http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

	} else {
		http.Redirect(w, r, "/register", http.StatusFound)
	}

}
