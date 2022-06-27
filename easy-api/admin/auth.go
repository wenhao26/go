package admin

import "github.com/gin-gonic/gin"

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")



}
