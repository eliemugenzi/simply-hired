package serializer

import (
	"fmt"
	"time"

	"github.com/eliemugenzi/simply-hired/db/models"
)

type JobResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	MinimumSalary uint64 `json:"min_salary"`
	MaximumSalary uint64 `json:"max_salary"`
	CreatedAt time.Time `json:"submitted_at"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	UserId uint `json:"user_id"`

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
	fmt.Println("USER ID", serializer.Job.UserId)
	return JobResponse{
		ID: serializer.Job.ID,
		Title: serializer.Job.Title,
		Description: serializer.Job.Description,
		MinimumSalary: serializer.Job.MinimumSalary,
		MaximumSalary: serializer.Job.MaximumSalary,
		ApplicationDeadline: serializer.Job.ApplicationDeadline,
		CreatedAt: serializer.Job.CreatedAt,
		UserId: serializer.Job.UserId,
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
