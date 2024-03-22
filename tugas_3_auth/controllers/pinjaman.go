package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"tugas_3_auth/database"
	"tugas_3_auth/helpers"
	"tugas_3_auth/models"
)

var (
	appJSON = "application/json"
)

func CreatePinjaman(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	pinjamanRequest := models.CreatePinjamanRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&pinjamanRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&pinjamanRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	pinjaman := models.Pinjaman{
		BukuId: pinjamanRequest.BukuId,
		UserId: userID,
	}

	// update status buku
	bukuId := pinjaman.BukuId
	db.Where("id = ?", bukuId).First(&models.Buku{}).Update("status", "Dipinjam")

	err := db.Debug().Create(&pinjaman).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	bukuString, _ := json.Marshal(pinjaman)
	bukuResponse := []models.GetPinjamanResponse{}
	json.Unmarshal(bukuString, &bukuResponse)

	c.JSON(http.StatusOK, bukuResponse)

	// c.JSON(http.StatusCreated, gin.H{
	// 	"id":             pinjaman.ID,
	// 	"tgl_peminjaman": pinjaman.CreatedAt,
	// 	"detail_buku": gin.H{
	// 		"buku_id":  pinjaman.BukuId,
	// 		"judul":    pinjaman.Buku.Judul,
	// 		"penulis":  pinjaman.Buku.Penulis,
	// 		"terbitan": pinjaman.Buku.Tahun,
	// 	},
	// 	"status": pinjaman.Buku.Status,
	// })
}

func GetPinjaman(c *gin.Context) {
	db := database.GetDB()
	pinjaman := []models.Pinjaman{}

	err := db.Debug().Preload("User").Preload("Buku").Order("id asc").Find(&pinjaman).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	pinjamanString, _ := json.Marshal(pinjaman)
	pinjamanResponse := []models.GetPinjamanResponse{}
	json.Unmarshal(pinjamanString, &pinjamanResponse)

	c.JSON(http.StatusOK, pinjamanResponse)
}

func DeletePinjaman(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	pinjamanId, _ := strconv.Atoi(c.Param("pinjamanId"))
	userID := uint(userData["id"].(float64))

	pinjaman := models.Pinjaman{}
	pinjaman.ID = uint(pinjamanId)
	pinjaman.UserId = userID

	err := db.Delete(&pinjaman).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// update status buku
	bukuId := pinjaman.BukuId
	db.Where("id = ?", bukuId).First(&models.Buku{}).Update("status", "Tersedia")

	c.JSON(http.StatusOK, gin.H{
		"message": "Buku sudah dikembalikan",
	})
}
