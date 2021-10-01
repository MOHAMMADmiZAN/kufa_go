package Route

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"kufa/Controllers"
	"kufa/DataBase"
	"net/http"
)

const PortNumber = ":5050"

func Route() {
	router := httprouter.New()
	//fs := http.FileServer(http.Dir("./assets"))
	//http.Handle("/resource/", http.StripPrefix("/resource/", fs))
	router.ServeFiles("/resource/*filepath", http.Dir("./assets"))
	router.GET("/", Controllers.Home)
	//http.HandleFunc("/blog", Controllers.Blog)
	router.GET("/blog", Controllers.Blog)
	//http.HandleFunc("/portfolioSingle", Controllers.PortfolioSingle)
	router.GET("/portfolioSingle", Controllers.PortfolioSingle)
	//http.HandleFunc("/blogDetails", Controllers.BlogDetails)
	router.GET("/blogDetails", Controllers.BlogDetails)
	//http.HandleFunc("/login", Controllers.Login)
	router.GET("/login", Controllers.Login)
	//http.HandleFunc("/register", Controllers.Register)
	router.GET("/register", Controllers.Register)
	//http.HandleFunc("/registerRequest", Controllers.RegisterRequest)
	router.POST("/registerRequest", Controllers.RegisterRequest)
	//http.HandleFunc("/loginRequest", Controllers.LoginRequest)
	router.POST("/loginRequest", Controllers.LoginRequest)
	router.GET("/logOut", Controllers.LogOut)
	//http.HandleFunc("/dashboard", Controllers.Dashboard)
	router.GET("/dashboard", Controllers.Dashboard)
	fmt.Println(fmt.Sprintf("Server Starting At PortNumber: %s", PortNumber))

	DataBase.Err = http.ListenAndServe(PortNumber, router)
	if DataBase.Err != nil {
		fmt.Println(DataBase.Err.Error())
	}

}
