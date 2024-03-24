package repositories

import (
	"log"

	"tugas_akhir/database"
	"tugas_akhir/models"
)

func RegisterRepository(user *models.User) error {
	db := database.GetDB()
	err := db.Create(&user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func LoginRepository(user *models.UserLogin) (*models.UserLogin, error) {
	db := database.GetDB()
	var userLogin models.UserLogin
	if err := db.Table("users").Where("users.username = ?", user.Username).First(&userLogin).Error; err != nil {
		log.Printf("Error while retrieving data user with err: %s", err)
		return nil, err
	}
	return &userLogin, nil
}

func UpdateUserRepository(Id int, user *models.UserUpdate) (*models.User, error) {
	db := database.GetDB()
	userData := models.User{}
	err := db.Preload("UserRole").Select("id", "username", "nama", "email", "alamat", "no_telp", "user_role_id", "created_at", "updated_at").First(&userData, Id)
	if err.Error != nil {
		log.Printf("Error get data user with err: %s", err.Error)
		return nil, err.Error
	}

	err = db.Model(&userData).Updates(
		models.User{
			Username: user.Username,
			Nama:     user.Nama,
			Email:    user.Email,
			Alamat:   user.Alamat,
			NoTelp:   user.NoTelp,
		})
	if err.Error != nil {
		log.Printf("Error update data user with err: %s", err.Error)
		return nil, err.Error
	}

	return &userData, nil
}

func DeleteUserRepository(Id int) error {
	db := database.GetDB()
	userData := models.User{}
	err := db.Unscoped().Delete(userData, Id)
	if err.Error != nil {
		log.Printf("Error delete data user with err: %s", err.Error)
		return err.Error
	}
	return nil
}
