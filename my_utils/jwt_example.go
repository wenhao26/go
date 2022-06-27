package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type claims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

const (
	SECRET = "12345" // 私钥
)

func genToken(data string) string {
	setClaims := claims{
		data,
		jwt.StandardClaims{
			Audience:  "接收JWT者",
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Id:        "唯一标识符",
			Issuer:    "发布者",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := reqClaims.SignedString([]byte(SECRET))
	if err != nil {
		return ""
	}

	return token
}

func checkToken(token string) (*claims, interface{}) {
	var c claims

	setToken, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不正确")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token已过期")
			} else {
				return nil, fmt.Errorf("token格式错误")
			}
		}
	}

	if setToken != nil {
		if key, ok := setToken.Claims.(*claims); ok && setToken.Valid {
			return key, nil
		} else {
			return nil, fmt.Errorf("token不正确")
		}
	}

	return nil, fmt.Errorf("token不正确")
}

func main() {
	token := genToken("自定义加密数据")
	fmt.Println("创建TOKEN：", token)

	ret, err := checkToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("校验TOKEN结果：", ret, ret.Data)
}
