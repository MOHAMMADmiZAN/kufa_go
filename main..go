package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/blog", blog)
	http.HandleFunc("/portfolio-single", portfolio_single)
	http.HandleFunc("/blog-details", blog_details)
	err := http.ListenAndServe(":9999", nil)
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
func portfolio_single(w http.ResponseWriter, r *http.Request) {
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
func blog_details(w http.ResponseWriter, r *http.Request) {
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
