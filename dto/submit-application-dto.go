package dto

type SubmitApplication struct {
  JobId uint `json:"job_id" binding:"required"`
  Body string `json:"body" binding:"required,min=20"`
}