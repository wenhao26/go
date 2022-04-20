package main

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



}
