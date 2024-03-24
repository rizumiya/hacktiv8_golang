package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	VehicleId          uint `json:"vehicle_id" validate:"required"`
	Vehicle            *Vehicle
	UserId             uint `json:"user_id"`
	User               *User
	ReportStatusId     uint `json:"status_laporan,omitempty"`
	ReportStatus       *ReportStatus
	TanggalPencurian   string `json:"tanggal_pencurian" validate:"required"`
	LokasiPencurian    string `json:"lokasi_pencurian" validate:"required"`
	KronologiPencurian string `json:"kronologi" validate:"required"`
}

/*
{
	"vehicle_id": 1,
	"user_id": 1,
	"status_laporan": 1,
	"tanggal_pencurian": "2006-01-02 3:04PM",
	"lokasi_pencurian": "di suatu gang",
	"kronologi": "kebetulan saya tinggal motor saya di gang dengan keadaan menyala"
}
*/

type ReportUpdate struct {
	TanggalPencurian   string `json:"tanggal_pencurian" validate:"required"`
	LokasiPencurian    string `json:"lokasi_pencurian" validate:"required"`
	KronologiPencurian string `json:"kronologi" validate:"required"`
}

type AdminReportUpdate struct {
	ReportStatusId     uint `json:"status_laporan" validate:"required"`
}
