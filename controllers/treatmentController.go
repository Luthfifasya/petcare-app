package controllers

import (
	"net/http"
	"petcare-app/database"
	"petcare-app/models"

	"github.com/gin-gonic/gin"
)

// CreateTreatment - tambah treatment baru
func CreateTreatment(c *gin.Context) {
	var treatment models.Treatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&treatment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create treatment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Treatment created successfully", "treatment": treatment})
}

// GetTreatments - ambil semua treatment
func GetTreatments(c *gin.Context) {
	var treatments []models.Treatment
	if err := database.DB.Find(&treatments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch treatments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"treatments": treatments})
}

// GetTreatmentByID - ambil treatment berdasarkan ID
func GetTreatmentByID(c *gin.Context) {
	id := c.Param("id")
	var treatment models.Treatment
	if err := database.DB.First(&treatment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"treatment": treatment})
}

// UpdateTreatment - update treatment
func UpdateTreatment(c *gin.Context) {
	id := c.Param("id")
	var treatment models.Treatment

	if err := database.DB.First(&treatment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}

	var input models.Treatment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&treatment).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Treatment updated successfully", "treatment": treatment})
}

// DeleteTreatment - hapus treatment
func DeleteTreatment(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Treatment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete treatment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Treatment deleted successfully"})
}
