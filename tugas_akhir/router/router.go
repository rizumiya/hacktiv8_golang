package router

import (
	"github.com/gin-gonic/gin"

	"tugas_akhir/controllers"
	"tugas_akhir/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/user/register", controllers.UserRegister)
		auth.POST("/user/login", controllers.UserLogin)
		auth.POST("/admin/register", controllers.AdminRegister)
		auth.POST("/admin/login", controllers.UserLogin)
	}

	user := r.Group("/user")
	userRouter := user.Group("/").Use(middlewares.Authentication()).Use(middlewares.UserAuthentication()).Use(middlewares.UserAuthorization())
	{
		userRouter.PUT("/:userId", middlewares.UserManipulationAuthorization(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.UserManipulationAuthorization(), controllers.DeleteUser)
	}
	userVehicleRouter := user.Group("/vehicle").Use(middlewares.Authentication()).Use(middlewares.UserAuthentication()).Use(middlewares.UserAuthorization())
	{
		userVehicleRouter.GET("", controllers.UserGetVehicles)
		userVehicleRouter.POST("", controllers.CreateVehicle)
		userVehicleRouter.PUT("/:vehicleId", middlewares.VehicleManipulationAuthorization(), controllers.UpdateVehicle)
		userVehicleRouter.DELETE("/:vehicleId", middlewares.VehicleManipulationAuthorization(), controllers.DeleteVehicle)
	}
	userReportRouter := user.Group("/report").Use(middlewares.Authentication()).Use(middlewares.UserAuthentication()).Use(middlewares.UserAuthorization())
	{
		userReportRouter.GET("", controllers.UserGetAllReports)
		userReportRouter.POST("", controllers.CreateReport)
		userReportRouter.PUT("/:reportId", middlewares.ReportManipulationAuthorization(), controllers.UpdateReport)
		userReportRouter.DELETE("/:reportId", middlewares.ReportManipulationAuthorization(), controllers.DeleteReport)
	}

	admin := r.Group("/admin")
	adminRouter := admin.Group("/").Use(middlewares.Authentication()).Use(middlewares.AdminAuthorization())
	{
		adminRouter.PUT("/:userId", middlewares.UserManipulationAuthorization(), controllers.UpdateUser)
		adminRouter.DELETE("/:userId", middlewares.UserManipulationAuthorization(), controllers.DeleteUser)
	}
	adminReportRouter := admin.Group("/report").Use(middlewares.Authentication()).Use(middlewares.AdminAuthorization())
	{
		adminReportRouter.GET("", controllers.AdminGetAllReports)
		adminReportRouter.PUT("/:reportId", controllers.AdminUpdateReportStatus)
	}

	return r
}
