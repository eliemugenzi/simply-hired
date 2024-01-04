package repository

import (
	"github.com/eliemugenzi/simply-hired/db/models"
	"gorm.io/gorm"
)

type JobRepo interface {
	SaveJob(job models.Job) (*gorm.DB, models.Job)
	FindJobsByRecruiter(userId uint) (*gorm.DB, []models.Job)
	GetSingleJob(jobId uint) (*gorm.DB, models.Job)
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

	result := repo.db.Where("user_id = ?", userId).Order("id desc").Find(&jobs).Preload("User")

	return result, jobs
}

func (repo *jobRepo) GetSingleJob(jobId uint) (*gorm.DB, models.Job) {
	job := models.Job{}
	result := repo.db.Where("id = ?", jobId).First(&job)
	return result, job
}