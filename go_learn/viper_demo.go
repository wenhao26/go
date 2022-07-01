package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// toml
	/*viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("redis.ip"))*/

	// env file
	/*viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
	}

	fmt.Println(viper.Get("APP_DEBUG"))
	fmt.Println(viper.Get("APP.NAME"))*/

	//viper.SetConfigFile("config.toml") // 指定配置文件路径
	//viper.SetConfigName("config") // 配置文件名称（无扩展名）
	//viper.SetConfigType("toml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")               // 还可以在工作目录中查找配置

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	/*// 在加载配置文件出错时
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
		} else {
			// 配置文件被找到，但产生了另外的错误
		}
	}*/

	fmt.Println(viper.GetString("app_name"))


}
