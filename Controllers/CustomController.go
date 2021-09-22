package Controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

func renderGohtml(w http.ResponseWriter, gohtml string) {

	parseFiles, err := template.ParseFiles("View/" + gohtml)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//var functions = template.FuncMap{}

//func renderGohtmlTest(w http.ResponseWriter) (map[string]*template.Template, error) {
//	myCache := map[string]*template.Template{}
//	pages, err := filepath.Glob("view/*.gohtml")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for _, page := range pages {
//		name := filepath.Base(page)
//		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
//		if err != nil {
//			return myCache, err
//		}
//		matchs, err := filepath.Glob("view/*.gohtml")
//		if err != nil {
//			return myCache, err
//		}
//		if len(matchs) > 0 {
//			ts, err = ts.ParseGlob("view/*.gohtml")
//			if err != nil {
//				return myCache, err
//			}
//
//		}
//		myCache[name] = ts
//	}
//	return myCache,nil
//
//}
func PasswordHash(password string) string {
	bs := []byte(password)
	hash, err2 := bcrypt.GenerateFromPassword(bs, bcrypt.DefaultCost)
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
	return string(hash)
}
func MatcherHash(hashPassword, password string) error {
	bs := []byte(password)
	hp := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hp, bs)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

type User struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=4,max=20"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}
type LoginUser struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=4,max=20" json:"password"`
}
