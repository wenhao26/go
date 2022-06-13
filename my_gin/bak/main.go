package main

import (
	"github.com/gin-gonic/gin"
	"go_project/go/my_gin/bak"
	"net/http"
)

var headers *bak.Headers

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers = bak.InitHeaders(c)
		//params.Header = InitHeaders(c)

		// log.Println(h)
		//CheckAccessToken(h.AccessToken)
		//c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare())

	{
		r.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "Index")
		})
		r.POST("/accessToken", func(c *gin.Context) {
			//CreateToken()

			c.String(http.StatusOK, "accessToken")
		})
	}

	r.Run(":8888")
}
