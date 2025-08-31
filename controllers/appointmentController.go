package controllers

import (
	"fmt"
	"net/http"
	"os"
	"petcare-app/database"
	"petcare-app/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

// CreateAppointment - buat appointment baru
func CreateAppointment(c *gin.Context) {
	var input struct {
		PetID       uint      `json:"pet_id"`
		TreatmentID uint      `json:"treatment_id"`
		Date        time.Time `json:"date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ambil treatment
	var treatment models.Treatment
	if err := database.DB.First(&treatment, input.TreatmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}

	// buat appointment
	appointment := models.Appointment{
		PetID:       input.PetID,
		TreatmentID: input.TreatmentID,
		Date:        input.Date,
		Status:      "scheduled",
	}
	if err := database.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	// ambil pet + user
	var pet models.Pet
	if err := database.DB.Preload("User").First(&pet, input.PetID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	// buat payment otomatis
	invoice := fmt.Sprintf("INV-%d-%s", appointment.ID, uuid.New().String())
	payment := models.Payment{
		AppointmentID: appointment.ID,
		InvoiceNumber: invoice,
		Amount:        treatment.Price,
		CustomerName:  pet.User.Name,
		CustomerEmail: pet.User.Email,
		Status:        "pending",
	}
	if err := database.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	// generate Midtrans Snap payment URL
	snapClient := snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.InvoiceNumber,
			GrossAmt: int64(payment.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: payment.CustomerName,
			Email: payment.CustomerEmail,
		},
	}

	snapResp, err := snapClient.CreateTransaction(snapReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payment.PaymentURL = snapResp.RedirectURL
	database.DB.Save(&payment)

	// reload appointment dengan semua relasi
	database.DB.
		Preload("Pet.User").
		Preload("Treatment").
		Preload("Payment").
		First(&appointment, appointment.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Appointment created successfully",
		"appointment": appointment,
	})
}

// GetAppointments - ambil semua appointment
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	if err := database.DB.Preload("Pet").Preload("Treatment").Preload("Payment").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch appointments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appointments": appointments})
}

// GetAppointmentByID - ambil appointment berdasarkan ID
func GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.Preload("Pet.User").Preload("Treatment").Preload("Payment").First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appointment": appointment})
}

// UpdateAppointment - update appointment
func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.Preload("Pet").Preload("Treatment").Preload("Payment").First(&appointment, id).Error; err != nil {
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
