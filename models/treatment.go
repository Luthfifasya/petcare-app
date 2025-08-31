package models

import "time"

// Treatment = layanan medis atau non-medis (grooming, vaksin, checkup, dll.)
type Treatment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Appointments []Appointment `json:"appointments"`
}
