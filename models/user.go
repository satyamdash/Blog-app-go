package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string  `json:"name"`
	Email *string `json:"email" gorm:"unique"`
	Posts []Post  `json:"posts"` // One-to-many relationship
}
