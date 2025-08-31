package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	AppointmentID uint           `json:"appointment_id"`
	InvoiceNumber string         `json:"invoice_number" gorm:"uniqueIndex"`
	Amount        float64        `json:"amount"`
	Status        string         `json:"status"`
	CustomerName  string         `json:"customer_name"`
	CustomerEmail string         `json:"customer_email"`
	PaymentURL    string         `json:"payment_url"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
