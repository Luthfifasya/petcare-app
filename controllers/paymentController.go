package controllers

import (
	"net/http"
	"petcare-app/database"
	"petcare-app/models"

	"github.com/gin-gonic/gin"
)

// MidtransCallback - handle callback dari Midtrans
func MidtransCallback(c *gin.Context) {
	var notif map[string]interface{}
	if err := c.ShouldBindJSON(&notif); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification"})
		return
	}

	orderID := notif["order_id"].(string)
	transactionStatus := notif["transaction_status"].(string)

	var payment models.Payment
	if err := database.DB.Where("invoice_number = ?", orderID).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Update status pembayaran
	payment.Status = transactionStatus
	database.DB.Save(&payment)

	c.JSON(http.StatusOK, gin.H{"message": "Payment status updated", "payment": payment})
}
