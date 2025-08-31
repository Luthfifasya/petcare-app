package main

<<<<<<< HEAD
=======
import (
	"fmt"
	"log"
	"os"

	"petcare-app/database"
	"petcare-app/routes"

	"github.com/gin-gonic/gin"
)

>>>>>>> f748c10 (finalisasi project (masih ada revisi sepertinya))
func main() {

<<<<<<< HEAD
<<<<<<< HEAD
=======
	// AutoMigrate
	database.DB.AutoMigrate(
		&models.User{},
		&models.Pet{},
		&models.Appointment{},
		&models.Treatment{},
	)

=======
>>>>>>> f748c10 (finalisasi project (masih ada revisi sepertinya))
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
>>>>>>> 57fa210 (menambahkan appointmentControoler)
}
