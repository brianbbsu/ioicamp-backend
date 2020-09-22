
package main

import (
	"github.com/gin-gonic/gin"
)

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

