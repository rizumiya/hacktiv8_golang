package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"tugas_2_api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root@tcp(127.0.0.1:3306)/order_by?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal terhubung ke database!")
		panic(err)
	}

	DB = database
	fmt.Println("Berhasil terhubung ke database!")
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Item{}, &models.Order{})

	if err != nil {
		fmt.Println("Gagal melakukan migrasi tabel!")
		panic(err)
	}

	fmt.Println("Berhasil melakukan migrasi tabel!")
}

func GetDb() *gorm.DB {
	return DB
}
