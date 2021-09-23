package Controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

func renderGohtml(w http.ResponseWriter, gohtml string, data ...interface{}) {
	parseFiles, err := template.ParseFiles("View/" + gohtml)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = parseFiles.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func renderMultipleGohtml(w http.ResponseWriter, temples []string, data ...interface{}) bool {
	parseFiles, err := template.ParseFiles(temples...)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//return false
		fmt.Println(err.Error())
	}
	err = parseFiles.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return true
}

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
