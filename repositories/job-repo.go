package repository

import (

	"github.com/eliemugenzi/simply-hired/db/models"
	"gorm.io/gorm"
)

type JobRepo interface {
	SaveJob(job models.Job) (*gorm.DB, models.Job)
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