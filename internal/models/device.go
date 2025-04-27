package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
}
