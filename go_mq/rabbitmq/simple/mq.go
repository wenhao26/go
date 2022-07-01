package simple

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	URL  string
	Conn *amqp.Connection
)

func init() {
	loadConfig()
	connect()
}

func loadConfig() {
	viper.SetConfigFile("simple/config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	URL = viper.GetString("url")
}

func connect() {
	var err error
	Conn, err = amqp.Dial(URL)
	if err != nil {
		panic(fmt.Errorf("Fatal error Connection: %s \n", err))
	}
}

func GetClient() *amqp.Connection {
	return Conn
}

func Channel(c *amqp.Connection) *amqp.Channel {
	ch, err := c.Channel()
	if err != nil {
		panic(fmt.Errorf("Fatal error Channel: %s \n", err))
	}
	return ch
}
