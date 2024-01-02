package models

import (
	"gorm.io/gorm"
)




type UserRole string

const (
	HR UserRole = "HR"
	APPLICANT UserRole = "APPLICANT"
)


type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role string `json:"role"`

	// The user association with the application
	Application Application
}
