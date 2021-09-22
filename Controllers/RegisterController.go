package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"kufa/DataBase"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	renderGohtml(w, "register.gohtml", nil)
}
func RegisterRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		confirmPassword := r.Form.Get("confirmPassword")
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
