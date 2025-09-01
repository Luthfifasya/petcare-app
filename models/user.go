package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Password  string    `json:"password"`
	Role      string    `json:"role" gorm:"default:customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Pets []Pet `json:"pets"` // Relasi One-to-Many
}
