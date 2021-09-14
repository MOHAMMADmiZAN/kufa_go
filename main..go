package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page/index.gohtml")
}
