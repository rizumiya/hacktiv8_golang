package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"tugas_akhir/database"
	"tugas_akhir/helpers"
	"tugas_akhir/models"
	"tugas_akhir/repositories"
	"tugas_akhir/validation"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.User

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if err := validation.ValidationFirst(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	user.Password = helpers.HashPass(user.Password)
	user.UserRoleId = 1 // user

	err := repositories.RegisterRepository(&user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	db := database.GetDB()
	db.Preload("UserRole").Where("id = ?", user.ID).First(&user)
	data := map[string]any{
		"_id": user.ID,
		"_info": gin.H{
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
		},
		"user": gin.H{
			"username":  user.Username,
			"nama":      user.Nama,
			"email":     user.Email,
			"alamat":    user.Alamat,
			"no_telp":   user.NoTelp,
			"user_role": user.UserRole.Role,
		},
	}

	helpers.NewHandlerResponse("New data user created", data).Success(c)
}

func AdminRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.User

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if err := validation.ValidationFirst(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	user.Password = helpers.HashPass(user.Password)
	user.UserRoleId = 2 // user

	err := repositories.RegisterRepository(&user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	db := database.GetDB()
	db.Preload("UserRole").Where("id = ?", user.ID).First(&user)
	data := map[string]any{
		"_id": user.ID,
		"_info": gin.H{
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"deleted_at": user.DeletedAt,
		},
		"user": gin.H{
			"username":  user.Username,
			"nama":      user.Nama,
			"email":     user.Email,
			"alamat":    user.Alamat,
			"no_telp":   user.NoTelp,
			"user_role": user.UserRole.Role,
		},
	}

	helpers.NewHandlerResponse("New data admin created", data).Success(c)
}

func UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.UserLogin

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if err := validation.ValidationFirst(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	userData, err := repositories.LoginRepository(&user)
	if err != nil {
		helpers.NewHandlerResponse("Username not found", nil).Failed(c)
		return
	}

	errHash := helpers.ComparePass([]byte(userData.Password), []byte(user.Password))
	if errHash != nil {
		log.Printf("Invalid Username or Password with err: %s\n", errHash)
		helpers.NewHandlerResponse("Username or password is incorrect", nil).BadRequest(c)
		return
	}

	token := helpers.GenerateToken(userData.ID, userData.Username)
	data := map[string]string{
		"token": token,
	}

	helpers.NewHandlerResponse("Successfully Login", data).Success(c)
}

func UpdateUser(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.UserUpdate

	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helpers.NewHandlerResponse("Error missing parameter", nil).BadRequest(c)
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if err := validation.ValidationFirst(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	userData, err := repositories.UpdateUserRepository(id, &user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	data := map[string]any{
		"_id": userData.ID,
		"_info": gin.H{
			"created_at": userData.CreatedAt,
			"updated_at": userData.UpdatedAt,
			"deleted_at": userData.DeletedAt,
		},
		"user": gin.H{
			"username":  userData.Username,
			"nama":      userData.Nama,
			"email":     userData.Email,
			"alamat":    userData.Alamat,
			"no_telp":   userData.NoTelp,
			"user_role": userData.UserRole.Role,
		},
	}
	helpers.NewHandlerResponse("Data updated successfully", data).Success(c)
}

func DeleteUser(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	var user models.User

	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helpers.NewHandlerResponse("Error missing parameter", nil).BadRequest(c)
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	if err := repositories.DeleteUserRepository(id); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("User deleted", nil).Success(c)
}
