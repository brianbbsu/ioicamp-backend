package main

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var defaultConfig = []byte(`
email:
  email: no-reply.ioicamp@csie.ntu.edu.tw
  password: owo
  server: smtp.csie.ntu.edu.tw
  port: 587
  requestDurationMinutes: 1
TemplateDir: ./template/
backend:
  port: 3030
  db: ./test.db
`)
// TODO: add verfication effective/expire time

// Config is the config for the whole package
var Config *viper.Viper

var db *gorm.DB

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
	// if err := sendEmailVerification("brianbb.su@gmail.com", token); err != nil {
	// 	panic(err)
	// }
	fmt.Println("Done!")

	initServer()
}

func initServer() {
	router := gin.Default()

	initRouter(router)
	initDatabase()

	router.Run("localhost:" + Config.GetString("backend.port"))
}

