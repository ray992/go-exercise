package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	id int64
	env string
	path string
	value string
}

func main() {

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
	var c config
	//查询单条
	queryErr := db.QueryRow(sqlStr, 1, "instamoney.va.fee").Scan(&c.id, &c.env, &c.path, &c.value)
	if queryErr != nil {
		fmt.Println("query err", queryErr)
	}
	fmt.Println(c.id, c.env, c.path, c.value)
	//多条查询
	moreSqlStr := "select id , env, path, value from config where id > ?"
	rows , queryErr := db.Query(moreSqlStr, 1)
	if queryErr != nil {
		fmt.Println("query more rows err", queryErr)
	}
	for rows.Next(){
		var c config
		forErr := rows.Scan(&c.id, &c.env, &c.path, &c.env)
		if forErr != nil {
			fmt.Println("for err")
			return
		}
		fmt.Println(c.id, c.env, c.path, c.value)
	}
	//插入
	/*insertStr := "insert into config(env, path, value) values (?, ?, ?)"
	insertRet, insertErr := db.Exec(insertStr, "dev", "wait.duration", 100)
	if insertErr != nil {
		fmt.Println("insert err", insertErr)
		return
	}
	curID, idErr := insertRet.LastInsertId()
	if idErr != nil {
		fmt.Println("id error", idErr)
	}
	fmt.Println(curID)*/
	//更新
	updateStr := "update config set value=? where id=?"
	updateRet, updateErr :=  db.Exec(updateStr, "1000", 35)
	if updateErr != nil {
		fmt.Println("update error", updateErr)
		return
	}
	affected, affectedErr := updateRet.RowsAffected()
	if affectedErr != nil {
		fmt.Println("affect error", affectedErr)
	}
	fmt.Println(affected)
	defer db.Close()
}
