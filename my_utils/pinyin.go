package main

import (
	"fmt"

	"github.com/mozillazg/go-pinyin"
)

func main() {
	c := "繁体中文"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(c, a))

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(c, a))

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(c, a))

}