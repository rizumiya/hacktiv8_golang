package controllers

import (
	// "log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "tugas_akhir/database"
	"tugas_akhir/helpers"
	"tugas_akhir/models"
	"tugas_akhir/repositories"
	"tugas_akhir/validation"
)

// admin
func AdminGetAllReports(c *gin.Context) {
	reportDatas, err := repositories.AdminGetAllReportRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data := []map[string]any{}
	for _, report := range *reportDatas {
		data = append(data, map[string]any{
			"_id": report.ID,
			"laporan": gin.H{
				"status":            report.ReportStatus.Status,
				"tanggal_pelaporan": report.CreatedAt,
				"tanggal_pencurian": report.TanggalPencurian,
				"lokasi_pencurian":  report.LokasiPencurian,
				"kronologi":         report.KronologiPencurian,
				"_pelapor": gin.H{
					"nama":    report.User.Nama,
					"email":   report.User.Email,
					"alamat":  report.User.Alamat,
					"no_telp": report.User.NoTelp,
				},
			},
			"motor": gin.H{
				"nomor_polisi":   report.Vehicle.NomorPolisi,
				"merk":           report.Vehicle.Merk,
				"tipe":           report.Vehicle.Tipe,
				"tahun_keluaran": report.Vehicle.TahunPembuatan,
				"warna":          report.Vehicle.WarnaKendaraan,
				"_pemilik": gin.H{
					"nama":   report.Vehicle.Pemilik,
					"kontak": report.Vehicle.NoTelpPemilik,
					"alamat": report.Vehicle.AlamatPemilik,
				},
				"_scan_surat": gin.H{
					"bpkb": report.Vehicle.ScanBpkb,
					"stnk": report.Vehicle.ScanStnk,
					"ktp":  report.Vehicle.ScanKtp,
				},
			},
		})
	}

	helpers.NewHandlerResponse("All reports retrieved successfully", data).Success(c)
}

func AdminUpdateReportStatus(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var report models.AdminReportUpdate

	id, err := strconv.Atoi(c.Param("reportId"))
	if err != nil {
		helpers.NewHandlerResponse("Error missing parameter", nil).BadRequest(c)
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&report)
	} else {
		c.ShouldBind(&report)
	}

	if err := validation.ValidationFirst(&report); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	reportData, err := repositories.AdminUpdateStatusRepository(id, &report)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	data := map[string]any{
		"_id": reportData.ID,
		"laporan": gin.H{
			"status":            reportData.ReportStatus.Status,
			"tanggal_pelaporan": reportData.CreatedAt,
			"tanggal_pencurian": reportData.TanggalPencurian,
			"lokasi_pencurian":  reportData.LokasiPencurian,
			"kronologi":         reportData.KronologiPencurian,
			"_pelapor": gin.H{
				"nama":    reportData.User.Nama,
				"email":   reportData.User.Email,
				"alamat":  reportData.User.Alamat,
				"no_telp": reportData.User.NoTelp,
			},
		},
		"motor": gin.H{
			"nomor_polisi":   reportData.Vehicle.NomorPolisi,
			"merk":           reportData.Vehicle.Merk,
			"tipe":           reportData.Vehicle.Tipe,
			"tahun_keluaran": reportData.Vehicle.TahunPembuatan,
			"warna":          reportData.Vehicle.WarnaKendaraan,
			"_pemilik": gin.H{
				"nama":   reportData.Vehicle.Pemilik,
				"kontak": reportData.Vehicle.NoTelpPemilik,
				"alamat": reportData.Vehicle.AlamatPemilik,
			},
			"_scan_surat": gin.H{
				"bpkb": reportData.Vehicle.ScanBpkb,
				"stnk": reportData.Vehicle.ScanStnk,
				"ktp":  reportData.Vehicle.ScanKtp,
			},
		},
	}

	helpers.NewHandlerResponse("Status updated successfully", data).Success(c)
}

// user
func UserGetAllReports(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	reportDatas, err := repositories.UserGetAllReportRepository(userID)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data := []map[string]any{}
	for _, report := range *reportDatas {
		data = append(data, map[string]any{
			"_id": report.ID,
			"laporan": gin.H{
				"status":            report.ReportStatus.Status,
				"tanggal_pelaporan": report.CreatedAt,
				"tanggal_pencurian": report.TanggalPencurian,
				"lokasi_pencurian":  report.LokasiPencurian,
				"kronologi":         report.KronologiPencurian,
			},
			"motor": gin.H{
				"nomor_polisi":   report.Vehicle.NomorPolisi,
				"merk":           report.Vehicle.Merk,
				"tipe":           report.Vehicle.Tipe,
				"tahun_keluaran": report.Vehicle.TahunPembuatan,
				"warna":          report.Vehicle.WarnaKendaraan,
				"_pemilik": gin.H{
					"nama":   report.Vehicle.Pemilik,
					"kontak": report.Vehicle.NoTelpPemilik,
					"alamat": report.Vehicle.AlamatPemilik,
				},
				"_scan_surat": gin.H{
					"bpkb": report.Vehicle.ScanBpkb,
					"stnk": report.Vehicle.ScanStnk,
					"ktp":  report.Vehicle.ScanKtp,
				},
			},
		})
	}

	helpers.NewHandlerResponse("All reports retrieved successfully", data).Success(c)
}

func CreateReport(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	var report models.Report

	if contentType == appJSON {
		c.ShouldBindJSON(&report)
	} else {
		c.ShouldBind(&report)
	}

	if err := validation.ValidationFirst(&report); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	reportData, err := repositories.CreateReportRepository(userID, &report)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	data := map[string]any{
		"_id": reportData.ID,
		"_info": gin.H{
			"created_at": reportData.CreatedAt,
			"updated_at": reportData.UpdatedAt,
			"deleted_at": reportData.DeletedAt,
		},
		"laporan": gin.H{
			"status":            reportData.ReportStatus.Status,
			"tanggal_pelaporan": reportData.CreatedAt,
			"tanggal_pencurian": reportData.TanggalPencurian,
			"lokasi_pencurian":  reportData.LokasiPencurian,
			"kronologi":         reportData.KronologiPencurian,
		},
		"motor": gin.H{
			"nomor_polisi":   reportData.Vehicle.NomorPolisi,
			"merk":           reportData.Vehicle.Merk,
			"tipe":           reportData.Vehicle.Tipe,
			"tahun_keluaran": reportData.Vehicle.TahunPembuatan,
			"warna":          reportData.Vehicle.WarnaKendaraan,
			"_pemilik": gin.H{
				"nama":   reportData.Vehicle.Pemilik,
				"kontak": reportData.Vehicle.NoTelpPemilik,
				"alamat": reportData.Vehicle.AlamatPemilik,
			},
			"_scan_surat": gin.H{
				"bpkb": reportData.Vehicle.ScanBpkb,
				"stnk": reportData.Vehicle.ScanStnk,
				"ktp":  reportData.Vehicle.ScanKtp,
			},
		},
	}

	helpers.NewHandlerResponse("New report added", data).SuccessCreate(c)
}

func UpdateReport(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var report models.ReportUpdate

	id, err := strconv.Atoi(c.Param("reportId"))
	if err != nil {
		helpers.NewHandlerResponse("Error missing parameter", nil).BadRequest(c)
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&report)
	} else {
		c.ShouldBind(&report)
	}

	if err := validation.ValidationFirst(&report); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	reportData, err := repositories.UpdateReportRepository(id, &report)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data := map[string]any{
		"_id": reportData.ID,
		"laporan": gin.H{
			"status":            reportData.ReportStatus.Status,
			"tanggal_pelaporan": reportData.CreatedAt,
			"tanggal_pencurian": reportData.TanggalPencurian,
			"lokasi_pencurian":  reportData.LokasiPencurian,
			"kronologi":         reportData.KronologiPencurian,
		},
		"motor": gin.H{
			"nomor_polisi":   reportData.Vehicle.NomorPolisi,
			"merk":           reportData.Vehicle.Merk,
			"tipe":           reportData.Vehicle.Tipe,
			"tahun_keluaran": reportData.Vehicle.TahunPembuatan,
			"warna":          reportData.Vehicle.WarnaKendaraan,
			"_pemilik": gin.H{
				"nama":   reportData.Vehicle.Pemilik,
				"kontak": reportData.Vehicle.NoTelpPemilik,
				"alamat": reportData.Vehicle.AlamatPemilik,
			},
			"_scan_surat": gin.H{
				"bpkb": reportData.Vehicle.ScanBpkb,
				"stnk": reportData.Vehicle.ScanStnk,
				"ktp":  reportData.Vehicle.ScanKtp,
			},
		},
	}

	helpers.NewHandlerResponse("Data updated successfully", data).Success(c)
}

func DeleteReport(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var report models.Report

	id, err := strconv.Atoi(c.Param("reportId"))
	if err != nil {
		helpers.NewHandlerResponse("Error missing parameter", nil).BadRequest(c)
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&report)
	} else {
		c.ShouldBind(&report)
	}

	if err := repositories.DeleteReportRepository(id); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Data deleted", nil).Success(c)
}
