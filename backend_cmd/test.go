package main

import (
	"fmt"
	"strings"

	"github.com/NaySoftware/go-fcm"

	"go_project/go/backend_cmd/service"
)

type FcmToken struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type FcmTopicSubscribe struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Type   int    `json:"type"`
	Topic  string `json:"topic"`
}

var (
	key  = "AAAAiKEVw4c:APA91bF_yS_l4ZGE3_DqJJP-vt1QOjuuWykub2PdkZxbgjcoiO__W3OMy5f03DIqfr8pYSrjXYkz-SyA85FQ05n-pDLOTbuFUVntIjfukaoeuAONNy3Pf4XghOxqsGediCmTWUuOGeQ1"
	lang = "TC"
	uid  = 10023
)

func langTopics(lang string, topics []FcmTopicSubscribe) {
	flags := [...]string{"CN", "EN", "TC", "ID", "TH", "VI", "PT", "RU"}

	for _, value := range topics {
		for _, f := range flags {
			if find := strings.Contains(value.Topic, f + "_"); find {
				fmt.Println(f, value.Topic)
				fmt.Println("find the character.")
			}
		}
	}
	fmt.Println(topics)
}

func main() {
	var ft FcmToken
	var fc []FcmTopicSubscribe

	db := service.MysqlClient()

	// 查询用户token
	res1 := db.Select("id,token").First(&ft, "user_id=?", uid)
	if res1.RowsAffected == 0 {
		fmt.Println("token尚未上报")
	}

	// 查询用户订阅的主题
	res2 := db.Select("id,user_id,type,topic").Find(&fc, "user_id=?", uid)
	if res2.RowsAffected == 0 {
		fmt.Println("没有订阅主题")
	}

	/*for key, value := range fc {
		fmt.Println(key,value)
	}*/

	// 组装语言订阅主题
	// langTopics(lang, fc)

	// 退订主题
	client := fcm.NewFcmClient(key)
	//client.BatchUnsubscribeFromTopic()
	fmt.Println(client)

}
