package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/gokufa")
	if err != nil {
		panic(err.Error())
	}
	TableNotExits()
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(0)
	Db.SetMaxIdleConns(0)
}
