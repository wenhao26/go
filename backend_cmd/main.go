package main

import (
	"flag"

	"go_project/go/backend_cmd/command"
)

func main() {
	cmd := flag.String("cmd", "", "Name of the execution command")
	flag.Parse()

	switch *cmd {
	case "fcm":
		command.FcmExecute()
		break
	}

}
