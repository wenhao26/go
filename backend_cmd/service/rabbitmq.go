package service

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	amqpUrl string
	err     error
	conn    *amqp.Connection
)

func init() {
	viper.SetConfigFile("config/config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	amqpUrl = viper.GetString("amqp_url")
}

func MQClient() *amqp.Connection {
	conn, err = amqp.Dial(amqpUrl)
	if err != nil {
		panic(fmt.Errorf("Fatal error Connection: %s \n", err))
	}
	return conn
}

func MQChannel(c *amqp.Connection) *amqp.Channel {
	ch, err := c.Channel()
	if err != nil {
		panic(fmt.Errorf("Fatal error Channel: %s \n", err))
	}
	return ch
}
