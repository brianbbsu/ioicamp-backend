
package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email string `json:"email" gorm:"not null"` // not null isnt working QQ
}
