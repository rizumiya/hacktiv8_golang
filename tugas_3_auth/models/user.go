package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;uniqueIndex"`
	Password string `json:"password" gorm:"not null"`
	Nama     string `json:"nama" gorm:"not null;uniqueIndex"`
}

type UserPinjamResponse struct {
	ID       uint   `json:"user_id"`
	Username string `json:"username"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nama     string `json:"nama"`
}
