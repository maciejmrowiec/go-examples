package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Listen      string
	Port        int
	Environment string
}

const (
	ConfigFileName = "config"
	NewConf        = `Listen: 127.0.0.1
Port: 8080
Enviroment: DEVELOPMENT
`
)

func main() {
	viper.New()

	// Config name
	viper.SetConfigName(ConfigFileName)

	// List of directoris where to search
	viper.AddConfigPath(".")

	// Reload config on config file modification
	// Uses filesystem notifications
	viper.WatchConfig()

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Print all settings - from original file
	fmt.Println(viper.AllSettings())

	// Overwrite config with new version
	if err := ioutil.WriteFile(ConfigFileName+".yml", []byte(NewConf), 0666); err != nil {
		log.Fatal(err)
	}

	// Wait for a moment
	time.Sleep(time.Second * 5)

	// Print all settings - after file modification
	// Sometimes I have seen issues connected to fsnotify library.
	// Viper uses version 1 by default, building viper with master seem to fix the issues,
	// in viper.go swap "gopkg.in/fsnotify.v1" for "github.com/go-fsnotify/fsnotify"
	fmt.Println(viper.AllSettings())
}
