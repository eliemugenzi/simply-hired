package repository

import (
	"fmt"

	"github.com/eliemugenzi/simply-hired/db/models"
	"gorm.io/gorm"
)

type JobRepo interface {
	SaveJob(job models.Job) (*gorm.DB, models.Job)
	FindJobsByRecruiter(userId uint) (*gorm.DB, []models.Job)
	GetSingleJob(jobId uint) (*gorm.DB, models.Job)
	SubmitApplication(applicationData models.Application) (*gorm.DB, models.Application)
}

type jobRepo struct {
	db *gorm.DB
}

func NewJobRepo(db *gorm.DB) *jobRepo {
	return &jobRepo{
		db: db,
	}
}

func (repo *jobRepo) SaveJob(job models.Job) (*gorm.DB, models.Job) {
    jobResult := repo.db.Create(&job)
	return jobResult, job
}

func (repo *jobRepo) FindJobsByRecruiter(userId uint) (*gorm.DB, []models.Job) {
	jobs := []models.Job{}

	result := repo.db.Where("user_id = ?", userId).Order("id desc").Preload("User").Find(&jobs).Preload("User")

	return result, jobs
}

func (repo *jobRepo) GetSingleJob(jobId uint) (*gorm.DB, models.Job) {
	job := models.Job{}
	result := repo.db.Where("id = ?", jobId).Preload("User").First(&job)
	fmt.Println(job, job.ToString())
	return result, job
}

func (repo *jobRepo) SubmitApplication(application models.Application) (*gorm.DB, models.Application) {
	applicationResult := repo.db.Create(&application)

	return applicationResult, application
}
