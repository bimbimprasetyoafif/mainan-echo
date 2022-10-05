package model

import "gorm.io/gorm"

type Lessons struct {
	gorm.Model
	Title string
	Users []Users
}
