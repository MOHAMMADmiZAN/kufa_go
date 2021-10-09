package Controllers

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"kufa/DataBase"
	"net/http"
)

type UsersDataTable struct {
	UserId     string
	UserEmail  string
	UserStatus string
	UserRole   string
}

func (u UsersDataTable) Add(i int) int {
	return i + 1
}

var temporaryAlert = SetFrontEndAlert("temp", "USER TEMPORARY DELETE", "temp-del")

func Users(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	userQuery := "SELECT id,email,status,role FROM `users` WHERE `status` = 1"
	usersDataRows, err := DataBase.Db.Query(userQuery)

	if err != nil {
		fmt.Println(err.Error())
	}

	var UsersData = UsersDataTable{}
	var UserDataResult []UsersDataTable

	for usersDataRows.Next() {
		var (
			id     string
			email  string
			status string
			role   string
		)
		err := usersDataRows.Scan(&id, &email, &status, &role)
		if err != nil {
			fmt.Println(err.Error())
		}
		UsersData.UserId = id
		UsersData.UserEmail = email
		UsersData.UserStatus = status
		UsersData.UserRole = role
		UserDataResult = append(UserDataResult, UsersData)

	}
	renderWithAuth(w, r, UserDataResult, "View/Dashboard/userList.gohtml")

}
func UserDelete(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	userId := r.URL.Query().Get("id")
	deletedStatus := "UPDATE `users` SET `status` = '%s' WHERE `id` = '%s'"
	deletedStatusSql := fmt.Sprintf(deletedStatus, "2", userId)
	deletedQuery, err := DataBase.Db.Query(deletedStatusSql)
	defer func(deletedQuery *sql.Rows) {
		err = deletedQuery.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(deletedQuery)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		temporaryAlert.AddAlertMessage(w, r)
		temporaryAlert.GetAlertMessage(w, r, nil, "View/Dashboard/userList.gohtml")
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

}
func UserTrash(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	userTrashQuery := "SELECT id,email,status,role FROM `users` WHERE `status` = 2"
	userTrashSql, err := DataBase.Db.Query(userTrashQuery)
	if err != nil {
		fmt.Println(err.Error())
	}
	userTrashData := UsersDataTable{}
	var userTrashResponse []UsersDataTable
	for userTrashSql.Next() {
		var (
			id     string
			email  string
			status string
			role   string
		)

		err := userTrashSql.Scan(&id, &email, &status, &role)
		if err != nil {
			fmt.Println(err.Error())
		}
		userTrashData.UserId = id
		userTrashData.UserEmail = email
		userTrashData.UserStatus = status
		userTrashData.UserRole = role
		userTrashResponse = append(userTrashResponse, userTrashData)
	}
	renderWithAuth(w, r, userTrashResponse, "View/Dashboard/recoverUser.gohtml")
}
func UserRecover(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	userId := r.URL.Query().Get("id")
	recoverQuery := "UPDATE `users` SET `status` = '%s' WHERE `id` = '%s'"
	recoverQuery = fmt.Sprintf(recoverQuery, "1", userId)
	recoverSql, err := DataBase.Db.Query(recoverQuery)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(recoverSql *sql.Rows) {
		err := recoverSql.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(recoverSql)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)

}
func UserPDelete(w http.ResponseWriter, r *http.Request, h httprouter.Params) {
	userId := r.URL.Query().Get("id")
	permanentDeletedQuery := "DELETE FROM `users` WHERE `id`='%s'"
	permanentDeletedQuery = fmt.Sprintf(permanentDeletedQuery, userId)
	permanentDeletedSql, _ := DataBase.Db.Query(permanentDeletedQuery)
	defer func(permanentDeletedSql *sql.Rows) {
		err := permanentDeletedSql.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(permanentDeletedSql)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}
