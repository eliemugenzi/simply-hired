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
	ID uint `gorm:"primaryKey,autoIncrement" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role string `json:"role"`
	Jobs []Job `gorm:"foreignKey:UserId"`
}
