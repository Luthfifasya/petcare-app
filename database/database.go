package database

import (
	"log"
	"os"

	"petcare-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env (optional untuk local)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in environment")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Database connected successfully")

	// Auto migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Pet{},
		&models.Appointment{},
		&models.Treatment{},
		&models.Payment{},
	)

	if err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	log.Println("✅ Database migrated successfully")
}
