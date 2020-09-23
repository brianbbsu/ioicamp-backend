
package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

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

func getEmailVerificationByEmailAndToken(email, token string) (EmailVerification, error) {
	var emailVerification EmailVerification

	result := db.Where("email = ? and token = ?", email, token).First(&emailVerification)

	return emailVerification, result.Error
}

func getLastCreatedAtByEmail(email string) (time.Time, error) {
	var response EmailVerification

	result := db.Select("created_at").Where("email = ?", email).Last(&response)

	if result.RowsAffected == 0 {
		return time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC), nil
	}

	return response.CreatedAt, result.Error
}

func getUserByEmailAndPassword(email, password string) (User, error) {
	var user User

	result := db.Where("email = ? and password = ?", email, password).First(&user)

	return user, result.Error
}

func createUserByEmailAndPassword(email, password string) (User, error) {
	user := User{Email: email, Password: password}

	result := db.Create(&user)

	return user, result.Error
}
