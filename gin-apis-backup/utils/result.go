package utils

import (
	"github.com/gin-gonic/gin"
	"go_project/go/gin-apis-backup/respond"
	"net/http"
	"time"
)

type Result struct {
	Time time.Time   `json:"time"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r := Result{}
	r.Time = time.Now()
	r.Code = int(respond.APICode.SUCCESS)
	r.Msg = respond.APICode.GetMessage(respond.APICode.SUCCESS)
	r.Data = data

	c.JSON(http.StatusOK, r)
}

func Error(c *gin.Context, code uint, msg string) {
	r := Result{}
	r.Time = time.Now()
	r.Code = int(code)
	//r.Msg = respond.APICode.GetMessage(code)
	r.Msg = msg
	r.Data = gin.H{}

	c.JSON(http.StatusOK, r)
}
