package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "kbe:kbe@tcp(192.168.0.103:3306)/fishtimer?parseTime=true")
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
