package main

import (
	_"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type App struct {
	Name   string `json:"name"`
	Type   int    `json:"type"`
	Source string `json:"source"`
}

func main() {
	s := time.Now()
	var app = new(App)

	app.Name = "CoinSky"
	app.Type = 1
	app.Source = "Overseas projects"

	// res, err := json.Marshal(app)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	res, err := json.Marshal(app)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	cost := time.Since(s)
	fmt.Printf("cost=[%s]", cost)
}
