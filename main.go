package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func main() {
	err:=initDB()
	if err!=nil{
		fmt.Println("init db failed , err: %v\n",err)
		return
	}

}
