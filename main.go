package main

import (
	"fmt"
	"log"
	"os"

	"petcare-app/database"
	"petcare-app/models"
	"petcare-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi database
	database.InitDB()

	// AutoMigrate
	database.DB.AutoMigrate(
		&models.User{},
		&models.Pet{},
		&models.Appointment{},
	)

	// Setup router
	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	fmt.Println("Server running on http://localhost" + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
