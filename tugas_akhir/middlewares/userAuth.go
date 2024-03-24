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

// Cek jika pemilik token memiliki role = user
func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		user := models.User{}

		db.First(&user, uint(userId))

		if user.UserRoleId != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Admins are not allowed to perform this action",
			})
			return
		}

		c.Next()
	}
}

// Cek jika pemilik token memiliki role = admin
func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		user := models.User{}

		db.First(&user, uint(userId))

		if user.UserRoleId != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Users are not allowed to perform this action",
			})
			return
		}

		c.Next()
	}
}

func UserManipulationAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		user := models.User{}

		err = db.Select("id").First(&user, uint(id)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "User doesn't exist",
			})
			return
		}

		if user.ID != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}
