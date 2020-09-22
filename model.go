
package main

import (
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email string `json:"email" gorm:"not null"` // not null isnt working QQ
}
// TODO: split request and database object

type EmailVerification struct {
	gorm.Model
	Email string `json:"email" gorm:"not null"`
	Token string `json:"token" gorm:"not null"`
}

func initDatabase() {
	var err error

	db, err = gorm.Open("sqlite3", Config.GetString("backend.db"))
	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&User{}, &EmailVerification{})
}

func getUserByEmail(email string) (User, error) {
	var user User

	result := db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func getEmailVerificationByEmail(email string) (EmailVerification, error) {
	var emailVerification EmailVerification

	result := db.Where("email = ?", email).First(&emailVerification)

	return emailVerification, result.Error
}
