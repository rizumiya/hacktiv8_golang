package middlewares

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"tugas_akhir/database"
	"tugas_akhir/models"
)

// Cek jika data vehicle yang ingin diolah adalah data yang dibuat oleh user itu sendiri
func VehicleManipulationAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		vehicleId, err := strconv.Atoi(c.Param("vehicleId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		vehicle := models.Vehicle{}

		err = db.Where("id = ?", uint(vehicleId)).First(&vehicle).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Vehicle not found",
			})
			return
		}

		err = db.Where("id = ?", uint(vehicleId)).Where("user_id = ?", userId).First(&vehicle).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You don't have permission for this action",
			})
			return
		}

		c.Next()
	}
}
