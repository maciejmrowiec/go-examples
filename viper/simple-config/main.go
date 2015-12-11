package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.New()

	// Config name
	viper.SetConfigName("config")

	// List of directoris where to search
	viper.AddConfigPath("/etc/viper-test")
	viper.AddConfigPath(".")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Print all settings
	fmt.Println(viper.AllSettings())
}
