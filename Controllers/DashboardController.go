package Controllers

import (
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var (
	Dkey   = []byte("super-secret-key")
	Dstore = sessions.NewCookieStore(Dkey)
)

func Dashboard(w http.ResponseWriter, r *http.Request, h httprouter.Params) {

	session, _ := Dstore.Get(r, "Go_Session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else {
		renderMultipleGohtml(w, nil, "View/Dashboard/dashboard.gohtml")
	}

}
