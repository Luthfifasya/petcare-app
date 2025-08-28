package models

import "time"

type Pet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Species   string    `json:"species"`
	Breed     string    `json:"breed"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`

	OwnerID uint `json:"owner_id"` // FK → User
	Owner   User `gorm:"foreignKey:OwnerID" json:"owner"`

	Appointments []Appointment `json:"appointments"`
}
