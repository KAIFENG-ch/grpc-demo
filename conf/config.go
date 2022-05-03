package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.AddConfigPath("./")
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(
		viper.GetString("mysql.Db"),
	)
}
