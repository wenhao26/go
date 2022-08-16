package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/coinsky_db?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect DB failed， err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(20)

}
