package common

import "log"

func ErrorMessage(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
