package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func main() {
	driverName := "root:MhxzKhl2020@tcp(127.0.0.1:3306)/cashpay"
	db, err :=  sqlx.Connect("mysql", driverName)
	if err != nil {
		 fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
}
