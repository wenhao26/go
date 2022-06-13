package routers

import (
	"github.com/gin-gonic/gin"
	"go_project/go/gin-apis-backup/handler"
)

func LoadRouter() {
	r := gin.Default()

	// 全局APIs令牌
	r.POST("/access-token", handler.AccessToken)

	_ = r.Run(":8888")
}
