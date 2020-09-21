package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
email:
  email: me@example.org
  password: thisisapassword
  server: smtp.csie.ntu.edu.tw
  port: 587
TemplateDir: ./template/
`)

// Config is the config for the whole package
var Config *viper.Viper

func initConfig() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.ReadConfig(bytes.NewBuffer(defaultConfig))
	err := Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %s", err))
	}
}

func main() {
	initConfig()
	token, err := getRandomToken(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(token, len(token))
	if err := sendEmailVerification("brianbb.su@gmail.com", token); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
