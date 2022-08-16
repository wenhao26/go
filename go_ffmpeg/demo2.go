package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		panic("Error1：" + err.Error())
	}
	fmt.Println(cmdPath)

	cmdArgs := []string{"-i D:\\go_project\\go\\go_ffmpeg"}
	//cmd := exec.Command(cmdPath, "-i")
	cmd := exec.Command(cmdPath, cmdArgs...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		panic("Error2：" + err.Error())
	}
	fmt.Println(string(result))
}
