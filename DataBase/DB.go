package DataBase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Db *sql.DB
var Err error

func init() {
	Db, Err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/gokufa")
	if Err != nil {
		panic(Err.Error())
	}
	TableNotExits()
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(0)
	Db.SetMaxIdleConns(0)
}
