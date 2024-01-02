package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
    UserId uint
	Body string `gorm:"type:varchar(255)" json:"body"`
	Status string `json:"status" binding:"oneof=PENDING IN_REVIEW REJECTED INTERVIEW ACCEPTED"`
}