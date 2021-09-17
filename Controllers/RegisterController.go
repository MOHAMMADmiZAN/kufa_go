package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"html/template"
	"kufa/View/DataBase"
	"log"
	"net/http"
)

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
			Email:           email,
			Password:        password,
			ConfirmPassword: confirmPassword,
		}
		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			fmt.Fprint(w, err.(validator.ValidationErrors))
		} else {
			hash := PasswordHash(password)
			insertQuery := "INSERT INTO users(email, password) VALUES ('%s','%s')"
			insertSql := fmt.Sprintf(insertQuery, email, hash)
			insert, err := DataBase.Db.Query(insertSql)
			if err != nil {
				log.Fatalln(err)
			}
			func(insert *sql.Rows) {
				err := insert.Close()
				if err != nil {
					log.Fatalln(err)
				}
			}(insert)
			fmt.Fprint(w, "Registration Successful")
		}

		//http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

	} else {
		http.Redirect(w, r, "/register", http.StatusFound)
	}

}
