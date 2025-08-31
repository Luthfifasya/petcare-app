package controllers

import (
	"net/http"
	"petcare-app/database"
	"petcare-app/models"

	"github.com/gin-gonic/gin"
)

// CreatePet - tambah pet
func CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi apakah user_id valid
	var user models.User
	if err := database.DB.First(&user, pet.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Simpan pet ke database
	if err := database.DB.Create(&pet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pet"})
		return
	}

	// Reload pet dengan relasi user
	if err := database.DB.Preload("User").Preload("Appointments").First(&pet, pet.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch created pet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pet created successfully",
		"pet":     pet,
	})
}

// GetPets - ambil semua pet
func GetPets(c *gin.Context) {
	var pets []models.Pet
	if err := database.DB.Preload("User").Find(&pets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pets": pets})
}

// GetPetByID - ambil pet by ID
func GetPetByID(c *gin.Context) {
	id := c.Param("id")
	var pet models.Pet
	if err := database.DB.Preload("User").First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pet": pet})
}

// UpdatePet - update data pet
func UpdatePet(c *gin.Context) {
	id := c.Param("id")
	var pet models.Pet
	if err := database.DB.Preload("User").First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	var input models.Pet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&pet).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Pet updated successfully", "pet": pet})
}

// DeletePet - hapus pet
func DeletePet(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Pet{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully"})
}
