
package main

import (
	// "log"

	"github.com/gin-gonic/gin"
)

func controller_getVerificationCode(c *gin.Context) {
	// TODO: make this function shorter
	var request EmailVerification

	c.BindJSON(&request)

	user, err := getUserByEmail(request.Email)

	if user.ID != 0 {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": "Email in use",
		})
		return
	}
	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": err,
		})
		return
	}

	email, err := getEmailVerificationByEmail(request.Email)

	// TODO: check last update

	// TODO: create new one
}

func controller_users_login(c *gin.Context) {
	var request, response User

	c.BindJSON(&request)

	result := db.Where(request).First(&response)
	// TODO: move this to `model.go`

	c.JSON(200, gin.H {
		"status": func() string { if result.RowsAffected > 0 && result.Error != nil { return "success" } else { return "failed" } } (),
		"user": response,
		"error": result.Error,
	})
}

func controller_users_register(c *gin.Context) {
	var request User

	c.BindJSON(&request)

	// TODO: check verification token

	result := db.Create(&request)
	// TODO: move this to `model.go`

	c.JSON(200, gin.H {
		"status": func() string { if result.RowsAffected > 0 && result.Error != nil { return "success" } else { return "failed" } } (),
		"user": request,
		"error": result.Error,
	})
}

// TODO: maybe need a error map?
