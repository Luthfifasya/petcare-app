package models

import "time"

type Appointment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PetID       uint      `json:"pet_id"` //FK
	TreatmentID uint      `json:"treatment_id"`
	ClinicID    uint      `json:"clinic_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`

	Pet       Pet       `gorm:"foreignKey:PetID" json:"pet"`
	Treatment Treatment `gorm:"foreignKey:TreatmentID" json:"treatment"`

	// Relasi One-to-One ke Payment
	Payment Payment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"payment"`
}
