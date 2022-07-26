package routers

import (
	"fmt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go_project/go/my_gin/controller"
)

func InitRouters() {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	adapter, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/gin_db?charset=utf8&parseTime=True&loc=Local", true)
	if err != nil {
		fmt.Println("连接数据库错误：", err.Error())
		return
	}

	e := casbin.NewEnforcer("model.conf", adapter)

	// 从DB加载策略
	e.LoadPolicy()

	r.GET("/", controller.Index)

	_ = r.Run(":8877")
}
