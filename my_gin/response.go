package main

import "encoding/json"

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(code int, message string, data interface{}) []byte {
	result := &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
	resp, _ := json.Marshal(result)
	return resp
}
