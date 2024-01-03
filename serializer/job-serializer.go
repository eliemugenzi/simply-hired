package serializer

import (
	"time"

	"github.com/eliemugenzi/simply-hired/db/models"
)

type JobResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	MinimumSalary uint64 `json:"min_salary"`
	MaximumSalary uint64 `json:"max_salary"`
	ApplicationDeadline time.Time `json:"application_deadline"`

	User UserResponse `json:"recruiter"`
}

type JobSerializer struct {
	Job models.Job
}

type JobsSerializer struct {
	Jobs []models.Job
}

func (serializer *JobSerializer) Response() JobResponse {
	userSerializer := UserSerializer{
		User: serializer.Job.User,
	}
	return JobResponse{
		ID: serializer.Job.ID,
		Title: serializer.Job.Title,
		Description: serializer.Job.Description,
		MinimumSalary: serializer.Job.MinimumSalary,
		MaximumSalary: serializer.Job.MaximumSalary,
		ApplicationDeadline: serializer.Job.ApplicationDeadline,
        User: userSerializer.Response(),
	}
}

func (serializer *JobsSerializer) Response() []JobResponse {
	response := []JobResponse{}
	for _, job := range serializer.Jobs {
		jobSerializer := JobSerializer {
			Job: job,
		}

		response = append(response, jobSerializer.Response())
	}

	return response
}
