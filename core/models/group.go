package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string
	AdminUserID uint
	InviteCode  string
	Lists       []List
}
