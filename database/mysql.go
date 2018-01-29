package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "rocky:3515666888_wcwzecePWX@tcp(45.77.43.2:3306)/kimmidoll?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	SqlDB.SetMaxIdleConns(200)
	SqlDB.SetMaxOpenConns(200)

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
