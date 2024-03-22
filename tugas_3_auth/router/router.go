package router

import (
	"github.com/gin-gonic/gin"

	"tugas_3_auth/controllers"
	"tugas_3_auth/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/", middlewares.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/", middlewares.Authentication(), controllers.DeleteUser)
	}

	bookRouter := r.Group("/buku")
	{
		bookRouter.GET("/", controllers.GetBuku)
	}

	bookProtectedRouter := r.Group("/admin/buku")
	{
		bookProtectedRouter.Use(middlewares.Authentication())
		bookProtectedRouter.POST("/", controllers.CreateBuku)
		bookProtectedRouter.DELETE("/:bukuId", middlewares.BukuAuthorization(), controllers.DeleteBuku)
	}

	pinjamanProtectedRouter := r.Group("/pinjaman")
	{
		pinjamanProtectedRouter.Use(middlewares.Authentication())
		pinjamanProtectedRouter.GET("/", controllers.GetPinjaman)
		pinjamanProtectedRouter.POST("/", controllers.CreatePinjaman)
		pinjamanProtectedRouter.DELETE("/:pinjamanId", middlewares.PinjamanAuthorization(), controllers.DeletePinjaman)
	}

	return r
}
