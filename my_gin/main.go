package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var headers *Headers

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers = InitHeaders(c)
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
