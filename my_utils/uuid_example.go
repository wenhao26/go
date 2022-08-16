package main

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {
	// 创建
	uv4, _ := uuid.NewV4()
	fmt.Println(uv4.String())

	// 解析
	parse, err := uuid.FromString(uv4.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(parse)
}
