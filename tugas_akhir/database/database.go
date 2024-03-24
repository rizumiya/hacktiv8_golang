package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"tugas_akhir/models"
)

var (
	dbInstance = "localhost:3306"
	username   = "root"
	dbName     = "laporan_pencurian"
	dbConnType = "tcp"
	db         *gorm.DB
)

func StartDB() {
	dsn := fmt.Sprintf("%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, dbConnType, dbInstance, dbName)
	sequel, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Cannot open sql")
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sequel,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("Cannot connect database")
	}
	defer fmt.Println("Successfully Connected to Database")

	db.Debug().AutoMigrate(models.UserRole{}, models.User{}, models.ReportStatus{}, models.Report{}, models.Vehicle{})

	CreateUserRole()
	CreateReportStatus()
}

func GetDB() *gorm.DB {
	return db
}

func CreateUserRole() {
	userRoles := []models.UserRole{
		{Role: "User"},
		{Role: "Admin"},
	}

	for _, userRole := range userRoles {
		var existingRole models.UserRole
		if err := db.Where("role = ?", userRole.Role).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&userRole).Error; err != nil {
					panic("failed to create userRole")
				}
				fmt.Printf("Role %s berhasil dibuat\n", userRole.Role)
			} else {
				panic("failed to query database")
			}
		} else {
			fmt.Printf("Role %s sudah ada\n", userRole.Role)
		}
	}
}

func CreateReportStatus() {
	reportStatuses := []models.ReportStatus{
		{Status: "Pending"},
		{Status: "Processing"},
		{Status: "Done"},
	}

	for _, reportStatus := range reportStatuses {
		var existingStatus models.ReportStatus
		if err := db.Where("status = ?", reportStatus.Status).First(&existingStatus).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&reportStatus).Error; err != nil {
					panic("failed to create reportStatus")
				}
				fmt.Printf("Status %s berhasil dibuat\n", reportStatus.Status)
			} else {
				panic("failed to query database")
			}
		} else {
			fmt.Printf("Status %s sudah ada\n", reportStatus.Status)
		}
	}
}
