package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"not null" json:"first_name"`
	Description string
	Done        bool
	UserId      uint
}
