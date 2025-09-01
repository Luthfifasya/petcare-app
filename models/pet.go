package models

import "time"

type Pet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Species   string    `json:"species"`
	Breed     string    `json:"breed"`
	Gender    string    `json:"gender"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Foreign key → User
	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`

	// Relasi One-to-Many → Appointment
	Appointments []Appointment `json:"appointments,omitempty"`
}
