package models

import (
	"time"

)

type Job struct {
	// gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	MinimumSalary uint64 `json:"min_salary"`
	MaximumSalary uint64 `json:"max_salary"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	Status string `json:"status"` // OPEN, CLOSED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId uint `json:"user_id"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserId;references:ID"`
}
