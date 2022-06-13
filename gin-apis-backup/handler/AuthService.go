package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_project/go/gin-apis-backup/respond"
	"go_project/go/gin-apis-backup/utils"
	"time"
)

type JWTClaims struct {
	AppId string `json:"app_id"`
	jwt.StandardClaims
}

const (
	TokenExpireDuration = time.Hour * 2
	Secret              = "sePIAPgVMJ23ahOBTRUzW14K9GHzKdQdgHX1eZkcHHVXv4Os7tFvGm25II8QbHtL"
)

func AccessToken(c *gin.Context) {
	appId := c.PostForm("app_id")
	token, err := GenToken(appId)
	if err != nil {
		utils.Error(c, respond.APICode.FAILED, err.Error())
		return
	}
	utils.Success(c, token)
	return
}

func GenToken(str string) (string, error) {
	claims := JWTClaims{
		str,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Secret))
}
