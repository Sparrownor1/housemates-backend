package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	PasswordHash string
	GroupID      *uint
	Group        Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
