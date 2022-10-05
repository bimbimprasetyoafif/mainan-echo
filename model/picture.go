package model

import (
	"gorm.io/gorm"
)

type ProfilePicture struct {
	gorm.Model
	Path    string `binding:"required"`
	UsersID uint
}
