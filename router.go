package main

import (
	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	routerInitAPIRouter(router.Group("/api"))
}

func routerInitAPIRouter(group *gin.RouterGroup) {
	routerAPIinitUserRouter(group.Group("/users"))

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	group.POST("get-verification-token", controllerGetVerificationCode)
	group.POST("login", controllerUsersLogin)
	group.POST("register", controllerUsersRegister)
}

func routerAPIinitUserRouter(group *gin.RouterGroup) {
	group.Use(authWithJWT)
	{
		group.GET("apply-form", controllerUsersGetApplyForm)
		group.PUT("apply-form", controllerUsersPutApplyForm)
		group.GET("whoami", controllerUsersWhoAmI)
	}
}
