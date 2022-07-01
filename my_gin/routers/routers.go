package routers

import (
	"github.com/gin-gonic/gin"
	"go_project/go/my_gin/controller"
)

func InitRouters() {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.Index)

	_ = r.Run(":8877")
}
