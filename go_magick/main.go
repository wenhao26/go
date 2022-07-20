package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Magick struct {
	Name string
}

var (
	mainInstruct = "magick"
)

// 识别图片信息
func ImgInfo(filename string) *Magick {
	return &Magick{
		Name: fmt.Sprintf("%s identify %s", mainInstruct, filename),
	}
}

// 执行命令
func (m *Magick) Execute() {
	//fmt.Println(m.Name)
	cmd := exec.Command("magick identify demo.jpg")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Complete...")
}

func main() {
	/*m := ImgInfo("demo.jpg")
	m.Execute()*/

	cmd := exec.Command("magick identify demo.jpg")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Complete...")
}
