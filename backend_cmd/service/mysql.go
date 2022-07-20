package service

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dsn string
	db  *gorm.DB
)

func init() {
	viper.SetConfigFile("config/config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	dsn = viper.GetString("mysql_dsn")
}

func MysqlClient() *gorm.DB {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cs_",
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}
