package main

import (
	"bytes"
	// "database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var defaultConfig = []byte(`
email:
  email: me@example.org
  password: thisisapassword
  server: smtp.csie.ntu.edu.tw
  port: 587
TemplateDir: ./template/
backend:
  port: 3030
  db: ./test.db
`)

// Config is the config for the whole package
var Config *viper.Viper

var userDb *gorm.DB

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

func initRouter(router *gin.Engine) {
	router_initApiRouter(router.Group("/api"))
}

func initDatabase() {
	var err error

	userDb, err = gorm.Open("sqlite3", Config.GetString("backend.db"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// userDb.Exec("create table users (username varchar(20) not null, password varchar(20) not null, email varchar(20) not null, primary key (username));")
	userDb.AutoMigrate(&User{})
}

// router

func router_initApiRouter(group *gin.RouterGroup) {
	router_api_initUserRouter(group.Group("/users"))

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
}

func router_api_initUserRouter(group *gin.RouterGroup) {
	group.POST("/login", controller_users_login)
	group.POST("/register", controller_users_register)
}

// controller

func controller_users_login(c *gin.Context) {
	var user User

	c.BindJSON(&user)

	result := userDb.Where(user).First(&user)

	c.JSON(200, gin.H {
		"status": func() string { if result.RowsAffected > 0 && result.Error != nil { return "success" } else { return "failed" } } (),
		"user": user,
		"error": result.Error,
	})
}

func controller_users_register(c *gin.Context) {
	var user User

	c.BindJSON(&user)

	result := userDb.Create(&user)

	c.JSON(200, gin.H {
		"status": func() string { if result.RowsAffected > 0 && result.Error != nil { return "success" } else { return "failed" } } (),
		"user": user,
		"error": result.Error,
	})
}

// model

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email string `json:"email" gorm:"not null"` // not null isnt working QQ
}
