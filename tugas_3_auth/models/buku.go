package models

import (
	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	UserId  uint `json:"user_id"`
	User    User
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"terbitan"`
	Status  string `json:"status"`
}

type BukuPinjamResponse struct {
	ID      uint   `json:"buku_id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"terbitan"`
	Status  string `json:"status"`
}

type CreateBukuRequest struct {
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"terbitan"`
}
