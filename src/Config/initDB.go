package Config

import (
	"database/sql"
)

var db *sql.DB

func initDB() (err error) {
	dns := "root:root@tcp(127.0.0.1:3306)/cyylog"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	err= db.Ping()
	if err!=nil{
		return err
	}
	defer db.Close()
	return err
}
