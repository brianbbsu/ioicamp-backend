package main

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User is a database model for an user account
type User struct {
	gorm.Model
	Email          string `gorm:"not null;unique"`
	HashedPassword []byte
	ApplyFormID    uint
	ApplyForm      ApplyForm `gorm:"not null"`
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
	ApplyFormData
}

// ApplyFormData contains application form fields
type ApplyFormData struct {
	// Email      string `gorm:"not null"`
	Name       string `json:"name" gorm:"not null"`
	Gender     string `json:"gender" gorm:"not null"`
	School     string `json:"school" gorm:"not null"`
	Grade      string `json:"grade" gorm:"not null"`
	CodeTime   string `json:"code-time" gorm:"not null"`
	CPTime     string `json:"cp-time" gorm:"not null"`
	Prize      string `json:"prize" gorm:"not null;size:1024"`
	OJ         string `json:"oj" gorm:"not null;size:1024"`
	Motivation string `json:"motivation" gorm:"not null;size:8000"`
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

func getUserByID(uid uint) (User, error) {
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
	user, err := getUserByEmail(email)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func createUserByEmailAndPassword(email, password string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err == nil {
		return user, errors.New("User already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), Config.GetInt("bcryptCost"))
	if err != nil {
		return user, errors.New("Unknown error")
	}
	user = User{Email: email, HashedPassword: hashedPassword}
	result := db.Create(&user)
	if result.Error != nil {
		return user, errors.New("Unknown error")
	}
	return user, nil
}

func attachApplyFormByUID(uid uint) error {
	var user User
	var form ApplyForm
	result := db.Create(&form)
	if result.Error != nil {
		return result.Error
	}
	db.First(&user, uid)
	user.ApplyFormID = form.ID
	db.Save(&user)
	return nil
}

func getApplyFormByUserID(uid uint) (ApplyForm, error) {
	var applyForm ApplyForm

	result := db.Find(&User{}, uid).Related(&applyForm)

	return applyForm, result.Error
}

func updateFormByForm(form ApplyForm) error {
	return db.Save(&form).Error
}

func updateUserByUser(user User) error {
	return db.Save(&user).Error
}
