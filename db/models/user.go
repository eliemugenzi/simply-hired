package models

import (
	"time"

)




type UserRole string

const (
	HR UserRole = "HR"
	APPLICANT UserRole = "APPLICANT"
)


type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role string `json:"role"`
	

	// Auto created records
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
