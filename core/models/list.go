package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name      string
	GroupID   uint
	ListItems []ListItem
}

type ListItem struct {
	gorm.Model
	Title  string
	Done   bool
	ListID uint
}
