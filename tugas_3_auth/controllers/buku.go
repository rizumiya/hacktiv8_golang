package controllers

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"tugas_3_auth/database"
	"tugas_3_auth/helpers"
	"tugas_3_auth/models"
)

func GetBuku(c *gin.Context) {
	db := database.GetDB()

	var buku []models.Buku

	if err := db.Preload("User").Find(&buku).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"buku": buku})
}

func CreateBuku(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	bukuRequest := models.CreateBukuRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&bukuRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&bukuRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	buku := models.Buku{
		Judul:   bukuRequest.Judul,
		Penulis: bukuRequest.Penulis,
		Tahun:   bukuRequest.Tahun,
		Status:  "Tersedia",
		UserId:  userID,
	}

	err := db.Debug().Create(&buku).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&buku, buku.ID).Error

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data berhasil ditambahkan",
	})
}

func DeleteBuku(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	bukuId, _ := strconv.Atoi(c.Param("bukuId"))
	userID := uint(userData["id"].(float64))

	buku := models.Buku{}
	buku.ID = uint(bukuId)
	buku.UserId = userID

	err := db.Delete(&buku).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus",
	})
}
