package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var wg sync.WaitGroup

type UserFcmToken struct {
	Id       int
	UserId   int
	Platform int
	Token    string
}

func main() {
	start := time.Now()
	dsn := "root:root@tcp(127.0.0.1:3306)/istory_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "iw_",
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go func() {
		var Results []UserFcmToken

		page := 0
		limit := 500
		for {
			/*if page == 3 {
				break
			}*/
			fmt.Println("Page=", page)
			offset := (page - 1) * limit
			//list := db.Debug().Limit(limit).Offset(offset).Find(&Results)
			list := db.Limit(limit).Offset(offset).Find(&Results)
			fmt.Println("Count：", list.RowsAffected)
			fmt.Println("Error：", errors.Is(list.Error, gorm.ErrRecordNotFound))

			for _, value := range Results {
				fmt.Println(value.Id, " => ", value.Token, "\n")
				/*fmt.Println(value.UserId)
				fmt.Println(value.Platform)
				fmt.Println(value.Token + "\n")*/
			}
			page++
		}

		defer wg.Done()
	}()
	wg.Wait()

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
