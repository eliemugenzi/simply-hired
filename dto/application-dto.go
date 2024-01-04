package dto

type Application struct {
  Body string `json:"body" binding:"required,min=20"`
}
