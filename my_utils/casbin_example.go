package main

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

func Auth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := "admin"
		fmt.Printf("obj=%s,act=%s,sub=%s", obj, act, sub)

		// 判断策略是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			c.String(http.StatusOK, "权限验证通过...\n")
			c.Next()
		} else {
			c.String(http.StatusOK, "抱歉！无权限操作...\n")
			c.Abort()
		}
	}
}

func main() {
	//+build ignore
	// go get github.com/casbin/gorm-adapter/v3
	// go get github.com/casbin/casbin/v2

	adapter, err := gormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/gin_db?charset=utf8&parseTime=True&loc=Local", true)
	if err != nil {
		fmt.Println("连接数据库错误：", err.Error())
	}

	fmt.Println(adapter)
	e, err := casbin.NewEnforcer("./to/model.conf", adapter)
	if err != nil {
		fmt.Println("初始化casbin错误：", err.Error())
	}

	// 从DB加载策略
	_ = e.LoadPolicy()

	r := gin.Default()

	// 获取
	r.GET("/get", func(c *gin.Context) {
		var data = make(map[int]string)
		c.String(http.StatusOK, "查看所有的策略...\n")
		list := e.GetPolicy()
		for _, item := range list {
			for k, v := range item {
				data[k] = v
			}
		}
		c.JSON(http.StatusOK, data)
	})

	// 增加
	r.GET("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "增加[/trade/list]策略...\n")
		if ok, _ := e.AddPolicy("admin", "/trade/list", "GET"); !ok {
			c.String(http.StatusOK, "[/trade/list]已存在...\n")
		} else {
			c.String(http.StatusOK, "增加[/trade/list]策略成功...\n")
		}
	})

	// 删除
	r.GET("/del", func(c *gin.Context) {
		c.String(http.StatusOK, "删除[trade/list]策略...\n")
		if ok, _ := e.RemovePolicy("admin", "/trade/list", "GET"); !ok {
			c.String(http.StatusOK, "[/trade/list]不存在或被删除...\n")
		} else {
			c.String(http.StatusOK, "删除[/trade/list]策略成功...\n")
		}
	})

	r.Use(Auth(e))
	r.GET("/trade/list", func(c *gin.Context) {
		c.String(http.StatusOK, "/trade/list\n")
	})
	r.GET("/trade/info", func(c *gin.Context) {
		c.String(http.StatusOK, "trade/info\n")
	})

	r.Run(":8877")
}
