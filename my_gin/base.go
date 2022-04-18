package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Headers struct {
	AppKey        string
	Version       string
	Timestamp     int
	Nonce         string
	Client        string
	Uuid          string
	AccessToken   string
	Authorization string
	Lang          string
	Timezone      int
	DeviceInfo    string
	Signature     string
}

func InitHeaders(c *gin.Context) *Headers {
	var headers Headers

	headers.AppKey = c.Request.Header.Get("App-Key")
	headers.Version = c.Request.Header.Get("App-Version")
	headers.Timestamp, _ = strconv.Atoi(c.Request.Header.Get("Timestamp"))
	headers.Nonce = c.Request.Header.Get("Nonce")
	headers.Client = c.Request.Header.Get("Client")
	headers.Uuid = c.Request.Header.Get("Uuid")
	headers.AccessToken = c.Request.Header.Get("Access-Token")
	headers.Authorization = c.Request.Header.Get("Authorization")
	headers.Lang = c.Request.Header.Get("Lang")
	headers.Timezone, _ = strconv.Atoi(c.Request.Header.Get("Timezone"))
	headers.DeviceInfo = c.Request.Header.Get("Device-Info")
	headers.Signature = c.Request.Header.Get("Signature")

	return &headers
}

func CheckAccessToken(token string) {
	fmt.Println("Data：", token)
}
