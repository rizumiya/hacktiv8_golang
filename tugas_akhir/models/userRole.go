package models

import (
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	Role string `json:"role_user"`
}
