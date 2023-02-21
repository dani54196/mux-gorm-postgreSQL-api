package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"not null"`
	Email    string `gorm:"not_null, unique_index"`
	Password string
	Tasks    []Task
}
