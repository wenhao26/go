package routers

import (
	"easy-api/admin"
	"github.com/gin-gonic/gin"
)

func InitAdminRouters() {
	r := gin.Default()

	r.POST("login", admin.Login)

	r.Run(":8888")
}
