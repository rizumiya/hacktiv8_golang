package repositories

import (
	"log"

	"tugas_akhir/database"
	"tugas_akhir/models"
)

func AdminGetAllReportRepository() (*[]models.Report, error) {
	db := database.GetDB()
	reports := []models.Report{}
	if err := db.Joins("User").Joins("ReportStatus").Joins("Vehicle").Model(&models.Report{}).Where("report_status_id IN ?", []int{1, 2}).Find(&reports).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &reports, nil
}

func AdminUpdateStatusRepository(Id int, report *models.AdminReportUpdate) (*models.Report, error) {
	db := database.GetDB()
	reportData := models.Report{}
	err := db.Preload("Vehicle").Preload("ReportStatus").First(&reportData, Id)
	if err.Error != nil {
		log.Printf("Error get data report with err: %s", err.Error)
		return nil, err.Error
	}

	err = db.Model(&reportData).Updates(
		models.Report{
			ReportStatusId: report.ReportStatusId,
		})
	if err.Error != nil {
		log.Printf("Error update status report with err: %s", err.Error)
		return nil, err.Error
	}
	
	if err := db.Preload("Vehicle").Preload("User").Preload("ReportStatus").First(&reportData, Id).Error; err != nil {
		return nil, err
	}
	return &reportData, nil
}

func UserGetAllReportRepository(Id uint) (*[]models.Report, error) {
	db := database.GetDB()
	reports := []models.Report{}
	if err := db.Joins("ReportStatus").Joins("Vehicle").Model(&models.Report{}).Where("reports.user_id = ?", Id).Find(&reports).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &reports, nil
}

func CreateReportRepository(Id uint, report *models.Report) (*models.Report, error) {
	db := database.GetDB()
	report.UserId = Id
	report.ReportStatusId = 1 // Pending
	err := db.Create(&report)
	if err.Error != nil {
		log.Printf("Error get data with err: %s", err.Error)
		return nil, err.Error
	}
	if err := db.Joins("ReportStatus").Joins("Vehicle").Model(&models.Report{}).First(&report, report.ID).Error; err != nil {
		return nil, err
	}
	return report, nil
}

func UpdateReportRepository(Id int, report *models.ReportUpdate) (*models.Report, error) {
	db := database.GetDB()
	reportData := models.Report{}
	err := db.Preload("Vehicle").First(&reportData, Id)
	if err.Error != nil {
		log.Printf("Error get data report with err: %s", err.Error)
		return nil, err.Error
	}

	err = db.Model(&reportData).Updates(
		models.Report{
			TanggalPencurian:   report.TanggalPencurian,
			LokasiPencurian:    report.LokasiPencurian,
			KronologiPencurian: report.KronologiPencurian,
		})
	if err.Error != nil {
		log.Printf("Error update data report with err: %s", err.Error)
		return nil, err.Error
	}

	if err := db.Joins("ReportStatus").Joins("Vehicle").Model(&models.Report{}).First(&reportData, Id).Error; err != nil {
		return nil, err
	}

	return &reportData, nil
}

func DeleteReportRepository(Id int) error {
	db := database.GetDB()
	userData := models.Report{}
	err := db.Unscoped().Delete(userData, Id)
	if err.Error != nil {
		log.Printf("Error delete data report with err: %s", err.Error)
		return err.Error
	}
	return nil
}
