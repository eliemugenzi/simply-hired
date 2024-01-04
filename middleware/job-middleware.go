package middleware

import (
	"net/http"
	"strconv"

	repository "github.com/eliemugenzi/simply-hired/repositories"
	"github.com/eliemugenzi/simply-hired/utils"
	"github.com/gin-gonic/gin"
)

type JobMiddleware interface {
	CheckJob() gin.HandlerFunc
}

type jobMiddleware struct {
	jobRepo repository.JobRepo
}

func NewJobMiddleware(jobRepo repository.JobRepo) *jobMiddleware {
	return &jobMiddleware{
		jobRepo: jobRepo,
	}
}

func (middleware *jobMiddleware) CheckJob() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jobId := ctx.Param("id")
		jobID, _ := strconv.ParseUint(jobId, 10, 10)

		_, job := middleware.jobRepo.GetSingleJob(uint(jobID))

		if job.ID == 0 {
			ctx.JSON(
				http.StatusNotFound,
				utils.GetResponse(
					http.StatusNotFound,
					"Job not found",
					nil,
				),
			)

			return
		}

		if job.Status != "OPEN" {
			ctx.JSON(
				http.StatusForbidden,
				utils.GetResponse(
					http.StatusForbidden,
					"This job is no longer open for new applications",
					nil,
				),
			)

			return
		}

		ctx.Set("job_id", job.ID)
	}
}