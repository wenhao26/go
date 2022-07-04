package main

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
)

const (
	//token = "fKJFdfP1TjOrjQ_FQsur5V:APA91bG6Xo8OcO_4E7s6WKpSxot4W1McsvUPyLRAmv87Miv4AD565A4JnYK7zyyBZd3WfAOfEbadWc3L4b4Hw5zqN3SzbeWSU9CCB9jHorrkI0kSivXCgzKHi41AGVl8vDA8E_87vmEO"
	token = "ep3A8C14RK6DD21SkHmaIo:APA91bEYC-tpb7T_I6tYBpcLHA9SPdEFmh6k1-NPyh5PYqY1Xxp0J8O9nNHklCH9YWAqBJOpyB1bYD6uln9zuBX3bKhO10YhQy0_rxbNZ5IpUNSDvGacROD7bMKDta4O_dgjQXeBW7Ci"
	key   = "AAAAiKEVw4c:APA91bF_yS_l4ZGE3_DqJJP-vt1QOjuuWykub2PdkZxbgjcoiO__W3OMy5f03DIqfr8pYSrjXYkz-SyA85FQ05n-pDLOTbuFUVntIjfukaoeuAONNy3Pf4XghOxqsGediCmTWUuOGeQ1"
)

func main() {
	data := map[string]string{
		"title": "鸡哥",
		"body":  "鸡哥",
		"msg":   "鸡哥",
	}

	ids := []string{
		token,
	}

	client := fcm.NewFcmClient(key)
	client.NewFcmRegIdsMsg(ids, data)
	status, err := client.Send()
	if err != nil {
		status.PrintResults()
	}
	fmt.Println(status)
}
