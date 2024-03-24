package controllers

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "tugas_akhir/database"
	"tugas_akhir/helpers"
	"tugas_akhir/models"
	"tugas_akhir/repositories"
	"tugas_akhir/validation"
)

func UserGetVehicles(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	vehicleDatas, err := repositories.UserGetVehiclesRepository(userId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data := []map[string]any{}
	for _, vehicle := range *vehicleDatas {
		data = append(data, map[string]any{
			"_id": vehicle.ID,
			"_info": gin.H{
				"created_at": vehicle.CreatedAt,
				"updated_at": vehicle.UpdatedAt,
				"deleted_at": vehicle.DeletedAt,
			},
			"motor": gin.H{
				"nomor_polisi":   vehicle.NomorPolisi,
				"merk":           vehicle.Merk,
				"tipe":           vehicle.Tipe,
				"tahun_keluaran": vehicle.TahunPembuatan,
				"warna":          vehicle.WarnaKendaraan,
				"_pemilik": gin.H{
					"nama":   vehicle.Pemilik,
					"kontak": vehicle.NoTelpPemilik,
					"alamat": vehicle.AlamatPemilik,
				},
				"_scan_surat": gin.H{
					"stnk": vehicle.ScanStnk,
					"ktp":  vehicle.ScanKtp,
					"bpkb": vehicle.ScanBpkb,
				},
			},
		})
	}

	helpers.NewHandlerResponse("All vehicles retrieved successfully", data).Success(c)
}

func CreateVehicle(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	var vehicle models.Vehicle

	if contentType == appJSON {
		c.ShouldBindJSON(&vehicle)
	} else {
		c.ShouldBind(&vehicle)
	}

	if err := validation.ValidationFirst(&vehicle); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	vehicleData, err := repositories.CreateVehicleRepository(userId, &vehicle)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	data := map[string]any{
		"_id": vehicleData.ID,
		"_info": gin.H{
			"created_at": vehicleData.CreatedAt,
			"updated_at": vehicleData.UpdatedAt,
			"deleted_at": vehicleData.DeletedAt,
		},
		"motor": gin.H{
			"nomor_polisi":   vehicleData.NomorPolisi,
			"merk":           vehicleData.Merk,
			"tipe":           vehicleData.Tipe,
			"tahun_keluaran": vehicleData.TahunPembuatan,
			"warna":          vehicleData.WarnaKendaraan,
			"_pemilik": gin.H{
				"nama":   vehicleData.Pemilik,
				"kontak": vehicleData.NoTelpPemilik,
				"alamat": vehicleData.AlamatPemilik,
			},
			"_scan_surat": gin.H{
				"stnk": vehicleData.ScanStnk,
				"ktp":  vehicleData.ScanKtp,
				"bpkb": vehicleData.ScanBpkb,
			},
		},
	}

	helpers.NewHandlerResponse("New vehicle added", data).SuccessCreate(c)
}

func UpdateVehicle(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var vehicle models.VehicleUpdate

	vehicleId, err := strconv.Atoi(c.Param("vehicleId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid parameter",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&vehicle)
	} else {
		c.ShouldBind(&vehicle)
	}

	if err := validation.ValidationFirst(&vehicle); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	vehicleData, err := repositories.UpdateVehicleRepository(uint(vehicleId), &vehicle)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data := map[string]any{
		"_id": vehicleData.ID,
		"_info": gin.H{
			"created_at": vehicleData.CreatedAt,
			"updated_at": vehicleData.UpdatedAt,
			"deleted_at": vehicleData.DeletedAt,
		},
		"motor": gin.H{
			"nomor_polisi":   vehicleData.NomorPolisi,
			"merk":           vehicleData.Merk,
			"tipe":           vehicleData.Tipe,
			"tahun_keluaran": vehicleData.TahunPembuatan,
			"warna":          vehicleData.WarnaKendaraan,
			"_pemilik": gin.H{
				"nama":   vehicleData.Pemilik,
				"kontak": vehicleData.NoTelpPemilik,
				"alamat": vehicleData.AlamatPemilik,
			},
			"_scan_surat": gin.H{
				"stnk": vehicleData.ScanStnk,
				"ktp":  vehicleData.ScanKtp,
				"bpkb": vehicleData.ScanBpkb,
			},
		},
	}

	helpers.NewHandlerResponse("Vehicle updated!", data).SuccessCreate(c)
}

func DeleteVehicle(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var vehicle models.Vehicle

	id, err := strconv.Atoi(c.Param("vehicleId"))
	if err != nil {
		helpers.NewHandlerResponse("Error missing parameter", nil).BadRequest(c)
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&vehicle)
	} else {
		c.ShouldBind(&vehicle)
	}

	if err := repositories.DeleteVehicleRepository(id); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Vehicle deleted", nil).Success(c)
}
