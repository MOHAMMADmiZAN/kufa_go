package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
)

//db  s *sql.DB
var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/gokufa")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(0)
}
func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/portfolio-single", portfolioSingle)
	http.HandleFunc("/blog-details", blogDetails)
	http.HandleFunc("/formRequest", formRequest)
	err = http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	parseFiles, err := template.ParseFiles("pages/index.gohtml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = parseFiles.Execute(w, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//http.ServeFile(w, r, "pages/index.gohtml")
}
func blog(w http.ResponseWriter, r *http.Request) {
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
	//http.ServeFile(w, r, "pages/index.gohtml")
}
func portfolioSingle(w http.ResponseWriter, r *http.Request) {
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
	//http.ServeFile(w, r, "pages/index.gohtml")
}
func blogDetails(w http.ResponseWriter, r *http.Request) {
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
	//http.ServeFile(w, r, "pages/index.gohtml")
}
func formRequest(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	insertQuery := "INSERT INTO `inbox` (`name`,`email`,`messages`) VALUES ('%s','%s','%s')"
	insertSql := fmt.Sprintf(insertQuery, name, email, message)
	insert, err := db.Query(insertSql)
	if err != nil {
		log.Fatalln(err)
	}
	func(insert *sql.Rows) {
		err := insert.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(insert)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	//fmt.Fprint(w, r.Header)
	//err := r.ParseForm()
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//form := r.Form
	//for i, val:= range form{
	//	_, err := fmt.Fprint(w,i,val)
	//	if err != nil {
	//		log.Fatalln(err.Error())
	//	}
	//}

}
