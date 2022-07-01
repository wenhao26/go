package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var Db *gorm.DB
var wg2 sync.WaitGroup

type BookTmp struct {
	BookId     int
	BookName   string
	BookNameCn string
	BookDesc   string
	Author     string
	CreateTime int
}

func init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/istory_db?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "iw_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return Db
}

func main() {
	db := GetDB()

	var bookTmp BookTmp

	t := time.Now()
	result1 := db.Debug().First(&bookTmp, "book_id=?", 1008)
	fmt.Println("Results：", result1)
	result2 := db.Debug().First(&bookTmp, "book_id=?", 1009)
	fmt.Println("Results：", result2)
	fmt.Println("普通方式耗时：", time.Since(t))

	t1 := time.Now()
	wg2.Add(2)
	go func() {
		result1 := db.Debug().First(&bookTmp, "book_id=?", 1008)
		fmt.Println("Results：", result1)
		defer wg2.Done()
	}()
	go func() {
		result1 := db.Debug().First(&bookTmp, "book_id=?", 1009)
		fmt.Println("Results：", result1)
		defer wg2.Done()
	}()
	wg2.Wait()
	fmt.Println("协程方式耗时：", time.Since(t1))

	fmt.Println("main end...")
}
