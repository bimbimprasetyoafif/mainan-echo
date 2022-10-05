package model

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	OrderID string
	QTY     int
}
