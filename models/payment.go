package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	AppointmentID uint      `json:"appointment_id"` // FK → Appointment
	Amount        float64   `json:"amount"`
	Method        string    `json:"method"` // QRIS, cash, transfer, etc.
	Status        string    `json:"status"` // paid, unpaid, refunded
	PaidAt        time.Time `json:"paid_at"`
}
