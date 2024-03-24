package repositories

import (
	"log"

	"tugas_akhir/database"
	"tugas_akhir/models"
)

func UserGetVehiclesRepository(Id uint) (*[]models.Vehicle, error) {
	db := database.GetDB()
	vehicles := []models.Vehicle{}
	if err := db.Where("user_id = ?", Id).Find(&vehicles).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &vehicles, nil
}

func CreateVehicleRepository(Id uint, vehicle *models.Vehicle) (*models.Vehicle, error) {
	db := database.GetDB()
	vehicle.UserId = Id
	err := db.Create(&vehicle)
	if err.Error != nil {
		log.Printf("Error get data with err: %s", err.Error)
		return nil, err.Error
	}
	if err := db.Where("id = ?", vehicle.ID).First(&vehicle).Error; err != nil {
		return nil, err
	}
	return vehicle, nil
}

func UpdateVehicleRepository(Id uint, vehicle *models.VehicleUpdate) (*models.Vehicle, error) {
	db := database.GetDB()
	vehicleData := models.Vehicle{}
	err := db.First(&vehicleData, Id)
	if err.Error != nil {
		log.Printf("Error get data vehicle with err: %s", err.Error)
		return nil, err.Error
	}

	err = db.Model(&vehicleData).Updates(
		models.Vehicle{
			NomorPolisi:    vehicle.NomorPolisi,
			Merk:           vehicle.Merk,
			Tipe:           vehicle.Tipe,
			TahunPembuatan: vehicle.TahunPembuatan,
			WarnaKendaraan: vehicle.WarnaKendaraan,
			Pemilik:        vehicle.Pemilik,
			AlamatPemilik:  vehicle.AlamatPemilik,
			NoTelpPemilik:  vehicle.NoTelpPemilik,
			ScanBpkb:       vehicle.ScanBpkb,
			ScanStnk:       vehicle.ScanStnk,
			ScanKtp:        vehicle.ScanKtp,
		})
	if err.Error != nil {
		log.Printf("Error update data vehicle with err: %s", err.Error)
		return nil, err.Error
	}

	return &vehicleData, nil
}

func DeleteVehicleRepository(Id int) error {
	db := database.GetDB()
	vehicleData := models.Vehicle{}
	err := db.Unscoped().Delete(vehicleData, Id)
	if err.Error != nil {
		log.Printf("Error delete data vehicle with err: %s", err.Error)
		return err.Error
	}
	return nil
}
