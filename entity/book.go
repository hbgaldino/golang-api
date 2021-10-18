package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint64
	Title       string
	Description string
}
