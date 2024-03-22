package main

import (
	"tugas_3_auth/database"
	"tugas_3_auth/router"
)

func main() {
	database.StartDB()

	r := router.StartApp()
	r.Run(":8080")
}
