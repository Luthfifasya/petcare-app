package models

import "time"

type Appointment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PetID       uint      `json:"pet_id"`
	TreatmentID uint      `json:"treatment_id"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`

	Pet       Pet       `gorm:"foreignKey:PetID" json:"pet"`
	Treatment Treatment `gorm:"foreignKey:TreatmentID" json:"treatment"`
	Payment   Payment   `gorm:"foreignKey:AppointmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payment"`
}
