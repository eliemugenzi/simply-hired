package dto

import "time"

type Job struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MinimumSalary uint64 `json:"min_salary" binding:"required"`
	MaximumSalary uint64 `json:"max_salary" binding:"required"`
	ApplicationDeadline time.Time `json:"application_deadline" binding:"required" time_format:"2006-01-02"`
}