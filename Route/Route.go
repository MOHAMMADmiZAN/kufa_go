package Route

import (
	"fmt"
	"kufa/Controllers"
	"kufa/DataBase"
	"net/http"
)

func Route() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	http.HandleFunc("/", Controllers.Home)
	http.HandleFunc("/blog", Controllers.Blog)
	http.HandleFunc("/portfolioSingle", Controllers.PortfolioSingle)
	http.HandleFunc("/blogDetails", Controllers.BlogDetails)
	http.HandleFunc("/login", Controllers.Login)
	http.HandleFunc("/register", Controllers.Register)
	http.HandleFunc("/registerRequest", Controllers.RegisterRequest)
	http.HandleFunc("/loginRequest", Controllers.LoginRequest)
	DataBase.Err = http.ListenAndServe(":9000", nil)
	if DataBase.Err != nil {
		fmt.Println(DataBase.Err.Error())
	}
}
