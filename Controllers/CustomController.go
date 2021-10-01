package Controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"kufa/KufaSessions"
	"log"
	"net/http"
	"path/filepath"
)

var (
	TemplatePath = "View/Dashboard/*.gohtml"
)

// single view render
func renderGohtml(w http.ResponseWriter, gohtml string, data interface{}) {
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

// LayoutFiles view path Catch
func LayoutFiles() []string {

	files, err := filepath.Glob(TemplatePath)
	if err != nil {
		panic(err)
	}
	return files
}

// multiple view render
func renderMultipleGohtml(w http.ResponseWriter, data interface{}, files ...string) {
	files = append(files, LayoutFiles()...)
	parseFiles, err := template.ParseFiles(files...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = parseFiles.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
	}

}

// log with Auth //
func renderWithAuth(w http.ResponseWriter, r *http.Request, data interface{}, files ...string) {
	session, _ := KufaSessions.Store.Get(r, "Go_Session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		renderMultipleGohtml(w, data, files...)

	}

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
		//fmt.Println(err.Error())
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

type FrontEndAlert struct {
	Type         string
	ErrorMessage string
	ErrorSession string
}

func (F FrontEndAlert) AddAlertMessage(w http.ResponseWriter, r *http.Request) {
	session, _ := KufaSessions.Store.Get(r, F.ErrorSession)
	session.AddFlash(F.ErrorMessage)
	err := session.Save(r, w)
	if err != nil {
		fmt.Println("Error Message Failed", err)
	}
	return
}
func (F FrontEndAlert) GetAlertMessage(w http.ResponseWriter, r *http.Request, renderGo ...string) {
	session, _ := KufaSessions.Store.Get(r, F.ErrorSession)
	if flashes := session.Flashes(); len(flashes) > 0 {
		ErrData := struct {
			Message interface{}
			Type    string
		}{
			Message: flashes[0],
			Type:    F.Type,
		}
		err := session.Save(r, w)
		if err != nil {
			fmt.Println("Error Message Failed", err)

		}
		renderMultipleGohtml(w, ErrData, renderGo...)
	}
}
