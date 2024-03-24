package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"unique" validate:"required"`
	Password   string `json:"password,omitempty" validate:"required,min=6"`
	Nama       string `json:"nama" validate:"required"`
	Email      string `json:"email" gorm:"unique" validate:"required,email"`
	Alamat     string `json:"alamat"`
	NoTelp     string `json:"no_telp"`
	UserRoleId uint   `json:"role_id"`
	UserRole   *UserRole

	Vehicles []Vehicle `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	Report   []Report  `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

/*
{
	"username": "admin 1",
	"password": "password 1",
	"nama": "ptby",
	"email": "asd@asd.asd",
	"alamat": "bumi",
	"no_telp": "123456789"
}
*/

type UserLogin struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Alamat   string `json:"alamat"`
	NoTelp   string `json:"no_telp"`
}
