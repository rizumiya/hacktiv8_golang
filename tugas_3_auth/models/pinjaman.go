package models

import (
	"gorm.io/gorm"
	"time"
)

type Pinjaman struct {
	gorm.Model
	UserId uint `json:"user_id"`
	User   User
	BukuId uint `json:"buku_id"`
	Buku   Buku
}

type CreatePinjamanRequest struct {
	BukuId uint `json:"buku_id"`
}

type GetPinjamanResponse struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	BukuId    uint      `json:"buku_id"`
	CreatedAt time.Time `json:"tgl_peminjaman"`
	User      UserPinjamResponse
	Buku      BukuPinjamResponse
}
