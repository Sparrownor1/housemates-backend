package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Name  string
	Items []ListItem
}
