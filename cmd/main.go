package main

import (
	"fmt"
	"log"

	"github.com/AlekseySauron/figures/app"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Ошибка viper", err.Error())
		log.Fatal("Ошибка viper", err)
		return
	}

	application := app.NewApplication()
	application.Run()
	defer application.Stop()
}
