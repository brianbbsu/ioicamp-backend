package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     Config.GetStringSlice("backend.allowedOrigin"),
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
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
	group.POST("get-password-reset-token", controllerGetPasswordResetToken)
	group.POST("password-reset", controllerPasswordReset)
}

func routerAPIinitUserRouter(group *gin.RouterGroup) {
	group.Use(authWithJWT)
	{
		group.GET("apply-form", controllerUsersGetApplyForm)
		group.PUT("apply-form", controllerUsersPutApplyForm)
		group.POST("change-password", controllerUsersChangePassword)
		group.GET("whoami", controllerUsersWhoAmI)
	}
}
