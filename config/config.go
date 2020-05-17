package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

var ServerConfig AppConfig

func InitConfig() *AppConfig {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	file, err := os.Open(pwd + "/config/config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := AppConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
