package service

import (

	"github.com/eliemugenzi/simply-hired/db/models"
	"github.com/eliemugenzi/simply-hired/dto"
	repository "github.com/eliemugenzi/simply-hired/repositories"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type JobService interface {
	SaveJob(jobDto dto.Job, userId uint) (*gorm.DB, models.Job)
	FindJobsByRecruiter(userId uint) (*gorm.DB, []models.Job)
	GetSingleJob(jobId uint) (*gorm.DB, models.Job)
}

type jobService struct {
	jobRepo repository.JobRepo
	authRepo repository.AuthRepo
}

func NewJobService(jobRepo repository.JobRepo, authRepo repository.AuthRepo) *jobService {
  return &jobService{
	jobRepo: jobRepo,
	authRepo: authRepo,
  }
}

func (service *jobService) SaveJob(jobDto dto.Job, userId uint) (*gorm.DB, models.Job) {
	jobModel := models.Job{}
	mapped := smapping.MapFields(&jobDto)

	err := smapping.FillStruct(&jobModel, mapped)

	if err != nil {
		panic(err)
	}

	_, user := service.authRepo.FindById(userId)

	jobModel.User = user
	jobModel.Status = "OPEN"

	return service.jobRepo.SaveJob(jobModel)
	
}

func (service *jobService) FindJobsByRecruiter(userId uint) (*gorm.DB, []models.Job) {
	result, jobs := service.jobRepo.FindJobsByRecruiter(userId)

	return result, jobs
}

func (service *jobService) GetSingleJob(jobId uint) (*gorm.DB, models.Job) {
	result, job := service.jobRepo.GetSingleJob(jobId)

	return result, job
}
