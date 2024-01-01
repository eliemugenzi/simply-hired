package utils

import (
	"github.com/eliemugenzi/simply-hired/dto"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


func InitializeCustomValidations() {
   if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	  _ = v.RegisterValidation("userRoleValidation", dto.CustomValidationHandler)
   }
}