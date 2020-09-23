package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// User is a database model for an user account
type User struct {
	gorm.Model
	Email       string `gorm:"not null;unique"`
	Password    string `gorm:"not null"`
	ApplyFormID int
	ApplyForm   ApplyForm `gorm:"not null"`
}

// EmailVerification is a database model storing email verification tokens
type EmailVerification struct {
	gorm.Model
	Email string `json:"email" gorm:"not null"`
	Token string `json:"token" gorm:"not null"`
}

// ApplyForm is a database model storing application forms
type ApplyForm struct {
	gorm.Model
	Email      string `gorm:"not null"`
	Name       string `gorm:"not null"`
	Gender     string `gorm:"not null"`
	School     string `gorm:"not null"`
	Grade      string `gorm:"not null"`
	CodeTime   string `gorm:"not null"`
	CPTime     string `gorm:"not null"`
	Prize      string `gorm:"not null;size:1024"`
	OJ         string `gorm:"not null;size:1024"`
	Motivation string `gorm:"not null;size:8000"`
}

func initDatabase() {
	var err error

	db, err = gorm.Open("sqlite3", Config.GetString("backend.db"))
	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&User{}, &EmailVerification{}, &ApplyForm{})
}

func getUserByEmail(email string) (User, error) {
	var user User

	result := db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func getUserByID(uid int) (User, error) {
	var user User

	result := db.First(&user, uid)

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

func getApplyFormByUserID(uid int) (ApplyForm, error) {
	var applyForm ApplyForm

	result := db.Find(&User{}, uid).Related(&applyForm)

	return applyForm, result.Error
}
