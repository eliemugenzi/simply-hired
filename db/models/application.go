package models

import "gorm.io/gorm"

/*
Application Status binding:"oneof=PENDING IN_REVIEW REJECTED INTERVIEW ACCEPTED"
*/

type Application struct {
	gorm.Model
    UserId uint `json:"user_id"`
	Body string `gorm:"type:varchar(255)" json:"body"`
	Status string `json:"status"`
	JobId uint `json:"job_id"`
	Job Job `json:"job" gorm:"foreignKey:JobId;references:ID"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserId;references:ID"`
}
