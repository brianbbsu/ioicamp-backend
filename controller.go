
package main

import (
	// "log"
	"time"

	"github.com/gin-gonic/gin"
)

func controller_getVerificationCode(c *gin.Context) {
	// TODO: make this function shorter
	var request EmailVerification

	c.BindJSON(&request)

	lastApplyTime, err := getLastCreatedAtByEmail(request.Email)
	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": err,
		})
		return
	}

	duration := time.Since(lastApplyTime)
	if duration.Minutes() < Config.GetFloat64("email.requestDurationMinutes") {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": "Request too fast",
		})
		return
	}

	token, err := getRandomToken(6)
	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": err,
		})
		return
	}

	result := db.Create(&EmailVerification{Email: request.Email, Token: token})
	if result.Error != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": err,
		})
		return
	}

	err = sendEmailVerification(request.Email, token)
	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H {
		"status": "success",
		"message": "email sent",
	})
}

func controller_users_login(c *gin.Context) {
	var request User

	c.BindJSON(&request)

	user, err := getUserByEmailAndPassword(request.Email, request.Password)

	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": "User not found",
		})
		return
	}

	c.JSON(200, gin.H {
		"status": "success",
		"user": user,
	})
}

func controller_users_register(c *gin.Context) {
	var request UserRegisterRequestInterface

	c.BindJSON(&request)

	emailVerification, err := getEmailVerificationByEmailAndToken(request.Email, request.Token)
	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": "Token not found",
		})
		return
	}

	duration := time.Since(emailVerification.CreatedAt)
	if duration.Minutes() > Config.GetFloat64("email.tokenEffectiveMinutes") {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": "Token expired",
		})
		return
	}

	// TODO: validate

	user, err := createUserByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		c.JSON(200, gin.H {
			"status": "failed",
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H {
		"status": "success",
		"user": user,
	})
}

type UserRegisterRequestInterface struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

// TODO: maybe need a error map?
