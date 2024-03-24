package middlewares

import (
	// "log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"tugas_akhir/database"
	"tugas_akhir/helpers"
	"tugas_akhir/models"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}

// Cek ke-valid-an token
func UserAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		userUsername := userData["username"]
		user := models.User{}

		if err := db.Select("id").First(&user, userId).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "User Not Found",
				"message": "Last chance to use your own account",
			})
			return
		}

		if err := db.Where("id = ?", userId).Where("username = ?", userUsername).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Unauthenticated",
				"message": "Please re-login using your current username and password",
			})
			return
		}
	}
}
