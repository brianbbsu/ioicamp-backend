package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/mattn/go-sqlite3"
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

var port int = 3030
var userDb *sql.DB

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

	initServer()
}

func initServer() {
	router := gin.Default()

	initRouter(router)
	initDatabase()

	router.Run(fmt.Sprint("localhost:", port))
}

func initRouter(router *gin.Engine) {
	router_initApiRouter(router.Group("/api"))
}

func initDatabase() {
	var err error

	userDb, err = sql.Open("sqlite3", "../test.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	userDb.Exec("create table users (username varchar(20) not null, password varchar(20) not null, email varchar(20) not null, primary key (username));")
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

}

// controller

func controller_users_login(c *gin.Context) {
	var user User

	c.BindJSON(&user)

	log.Println(user.Username, user.Password)
	stmt, err := userDb.Prepare("select * from users where username = ? and password = ?")

	if err != nil {
		log.Fatal(err)
		return
	}
	
	rows, err := stmt.Query(user.Username, user.Password)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(200, gin.H {
		"status":  "success",
		"username": user.Username,
		"password": user.Password,
		"email": user.Email,
	})
}

// model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}
