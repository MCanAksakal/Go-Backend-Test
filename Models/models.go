package Models

import "gorm.io/gorm"

type UserAuth struct {
	gorm.Model
	ID       int
	UserName string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Mail     string `gorm:"unique"`
	Password string
}
