package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name           string
	Email          string
	ProfilePicture ProfilePicture
}

type UsersCommon struct {
	Name  string
	Email string
}
