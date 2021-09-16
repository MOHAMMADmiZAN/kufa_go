package main

import (
	"fmt"
	"net/http"
)

func Route() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	http.HandleFunc("/", Home)
	http.HandleFunc("/blog", Blog)
	http.HandleFunc("/portfolioSingle", PortfolioSingle)
	http.HandleFunc("/blogDetails", BlogDetails)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/registerRequest", RegisterRequest)
	http.HandleFunc("/inboxRequest", InboxRequest)
	err = http.ListenAndServe(":5555", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
