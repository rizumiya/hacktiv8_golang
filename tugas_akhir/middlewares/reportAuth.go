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

func ReportManipulationAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		reportId, err := strconv.Atoi(c.Param("reportId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		report := models.Report{}

		err = db.Where("id = ?", uint(reportId)).First(&report).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "report not found",
			})
			return
		}

		err = db.Where("id = ?", uint(reportId)).Where("user_id = ?", userId).First(&report).Error
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
