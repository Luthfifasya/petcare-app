package controllers

import (
	"net/http"
	"petcare-app/database"
	"petcare-app/models"

	"github.com/gin-gonic/gin"
)

// CreateAppointment - buat appointment baru
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully", "appointment": appointment})
}

// GetAppointments - ambil semua appointment
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	if err := database.DB.Preload("Pet").Preload("Treatment").Preload("Treatment").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appointments": appointments})
}

// GetAppointmentByID - ambil appointment berdasarkan ID
func GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.Preload("Pet.User").Preload("Treatment").First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appointment": appointment})
}

// UpdateAppointment - update appointment
func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	var input models.Appointment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&appointment).Updates(input)
	c.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully", "appointment": appointment})
}

// DeleteAppointment - hapus appointment
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}
