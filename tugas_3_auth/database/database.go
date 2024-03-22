package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"tugas_3_auth/models"
)

var (
	dbInstance = "localhost:3306"
	username   = "root"
	dbName     = "test"
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

	db.Debug().AutoMigrate(models.User{}, models.Buku{}, models.Pinjaman{})

	// AddDataBuku()
}

func GetDB() *gorm.DB {
	return db
}

func AddDataBuku() {
	dataBuku := []models.Buku{
		{Judul: "Buku 1", Penulis: "Penulis 1", Tahun: 2023, Status: "Tersedia"},
		{Judul: "Buku 2", Penulis: "Penulis 2", Tahun: 2022, Status: "Tersedia"},
		{Judul: "Buku 3", Penulis: "Penulis 3", Tahun: 2021, Status: "Tersedia"},
	}

	var db = GetDB()

	for _, book := range dataBuku {
		var existingBook models.Buku
		if err := db.Where("judul = ?", book.Judul).First(&existingBook).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&book).Error; err != nil {
					panic("Gagal menambahkan buku")
				}
				fmt.Printf("Buku %s berhasil dibuat\n", book.Judul)
			} else {
				panic("query SQL salah..")
			}
		} else {
			fmt.Printf("Buku %s sudah ada\n", book.Judul)
		}
	}
}
