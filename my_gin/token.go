package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(uuid string) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  "https://test.me",
		ExpiresAt: time.Now().Unix() + int64(ACCESS_TOKEN_TTL),
		Id:        uuid,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "localhost",
		NotBefore: time.Now().Unix(),
		Subject:   "APIs",
	}
	var jwtSecret = []byte(APP_SECRET)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}
