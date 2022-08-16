package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmdPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(cmdPath)

	//cmd := exec.Command(cmdPath, "-version")
	cmdArgs := []string{"-i", "demo.mp4"}
	cmd := exec.Command(cmdPath, cmdArgs...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("command output: %q", out.String())

	/*cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println("执行错误：", err.Error())
	}*/

	/*result, err := cmd.CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(result))*/
}
