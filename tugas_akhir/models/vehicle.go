package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	UserId         uint `json:"user_id"`
	User           *User
	NomorPolisi    string `json:"nomor_polisi" gorm:"unique" validate:"required"`
	Merk           string `json:"merk" validate:"required"`
	Tipe           string `json:"tipe"`
	TahunPembuatan int    `json:"tahun_keluaran" validate:"required"`
	WarnaKendaraan string `json:"warna" validate:"required"`
	Pemilik        string `json:"pemilik" validate:"required"`
	AlamatPemilik  string `json:"alamat_pemilik" validate:"required"`
	NoTelpPemilik  string `json:"kontak_pemilik" validate:"required"`
	ScanBpkb       string `json:"scan_bpkb"`
	ScanStnk       string `json:"scan_stnk"`
	ScanKtp        string `json:"scan_ktp"`
}

/*
{
	"nomor_polisi": "AD2456VB",
	"merk": "Honda",
	"tipe": "bebek",
	"tahun_keluaran": 1992,
	"warna": "hitam",
	"pemilik": "agus",
	"alamat_pemilik": "Bumi_745",
	"kontak_pemilik": "+3276543565",
	"scan_bpkb": "img/scan/bpkb.jpg",
	"scan_stnk": "img/scan/stnk.jpg",
	"scan_ktp": "img/scan/ktp.jpg"
}
*/

type VehicleUpdate struct {
	NomorPolisi    string `json:"nomor_polisi" gorm:"unique"`
	Merk           string `json:"merk"`
	Tipe           string `json:"tipe"`
	TahunPembuatan int    `json:"tahun_keluaran"`
	WarnaKendaraan string `json:"warna"`
	Pemilik        string `json:"pemilik"`
	AlamatPemilik  string `json:"alamat_pemilik"`
	NoTelpPemilik  string `json:"kontak_pemilik"`
	ScanBpkb       string `json:"scan_bpkb"`
	ScanStnk       string `json:"scan_stnk"`
	ScanKtp        string `json:"scan_ktp"`
}
