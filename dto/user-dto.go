package dto

import (
	"github.com/eliemugenzi/simply-hired/db/models"
	"github.com/go-playground/validator/v10"
)

func StringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func CustomValidationHandler(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	s := []string {string(models.APPLICANT), string(models.HR)}

    return StringContains(s, input)
}

type User struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" binding:"oneof=APPLICANT HR"`
	/*
	* In case you want to add a custom validation handler
	*/
	// Role string `json:"role" binding:"userRoleValidation"`

}