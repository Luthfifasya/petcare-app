package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Password  string    `json:"-"` // disembunyikan dari JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi One-to-Many → Pet
	Pets []Pet `gorm:"foreignKey:UserID" json:"pets"`
}
