package command

import (
	"encoding/json"
	"fmt"
	"log"

	"go_project/go/backend_cmd/service"
)

type Message struct {
	Lang       string `json:"lang"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
	Platform   int    `json:"platform"`
	CreateTime int    `json:"create_time"`
}

type FcmToken struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func FcmExecute() {
	client := service.MQClient()
	ch := service.MQChannel(client)
	defer ch.Close()

	forever := make(chan bool)
	go func() {
		msgData, err := ch.Consume("fcm-token-queue", "", false, false, false, false, nil)
		if err != nil {
			fmt.Println(err)
		}

		var m Message
		var ft FcmToken

		for msg := range msgData {
			log.Println("Message：", string(msg.Body))
			_ = msg.Ack(false) // Ack

			_ = json.Unmarshal([]byte(msg.Body), &m)

			db := service.MysqlClient()
			res := db.Select("id,token").First(&ft, "user_id=?", m.UserId)
			if res.RowsAffected == 0 {
				log.Printf("Err：uid=%d，尚未上报token", m.UserId)
				continue
			}


			fmt.Println(ft)

		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
