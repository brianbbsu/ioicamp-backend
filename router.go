
package main

import (
	// "log"

	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	router_initApiRouter(router.Group("/api"))
}

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
