package models

import (
	"gorm.io/gorm"
)

type ReportStatus struct {
	gorm.Model
	Status string `json:"status"`
}
