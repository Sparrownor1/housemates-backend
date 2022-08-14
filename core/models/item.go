package models

import "gorm.io/gorm"

type ListItem struct {
	gorm.Model
	Title string
	Done  bool
}
