package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	
	"tugas_2_api/database"
	"tugas_2_api/handler"
)

func main() {

	database.ConnectDatabase()
	database.AutoMigrate()

	router := gin.Default()

	router.POST("/orders", handler.CreateOrder)
	router.GET("/orders", handler.GetOrders)
	router.PUT("/orders/:orderId", handler.UpdateOrder)
	router.DELETE("/orders/:orderId", handler.DeleteOrder)

	log.Fatal(http.ListenAndServe(":8080", router)) 

}
