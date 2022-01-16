package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


type config2 struct {
	id int64
	env string
	path string
	value string
}

var db *sql.DB

func main() {
	//预处理
	driverName := "root:MhxzKhl2020@tcp(127.0.0.1:3306)/cashpay"
	db, err := sql.Open("mysql", driverName)
	if err != nil {
		panic(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}
	db.SetMaxOpenConns(10) //设置数据库建立连接的最大数目
	db.SetMaxIdleConns(5) //设置最大闲置连接数
	sqlStr := "select id, env, path, value from config where id=? and path=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		 fmt.Println("prepare failed", err)
		 return
	}
	defer stmt.Close()
	rows , queryErr := stmt.Query(1, "instamoney.va.fee")
	if queryErr != nil {
		fmt.Println("stmt query failed", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var c1 config2
		err := rows.Scan(&c1.id, &c1.env, &c1.path, &c1.value)
		if	err != nil{
			continue
		}
		fmt.Println(c1.id, c1.env, c1.path, c1.value)
	}
	//事务
	tx , err := db.Begin()
	if err != nil {
		tx.Rollback()
		return
	}
	updateStr := "update config set value=? where id=?"
	result, updateErr :=  tx.Exec(updateStr, "1000", 35)
	if updateErr != nil {
		tx.Rollback()
		return
	}
	row , err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return
	}
	fmt.Println(row)
	tx.Commit()
}
