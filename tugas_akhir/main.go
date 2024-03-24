package main

import (
	"tugas_akhir/database"
	"tugas_akhir/router"
)

func main() {
	database.StartDB()

	r := router.StartApp()
	
	r.Run(":8080")
}
