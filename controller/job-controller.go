package controller

import (
	"net/http"
	"strconv"

	"github.com/eliemugenzi/simply-hired/dto"
	"github.com/eliemugenzi/simply-hired/serializer"
	service "github.com/eliemugenzi/simply-hired/services"
	"github.com/eliemugenzi/simply-hired/utils"
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
)

type JobController interface {
	SaveJob(context *gin.Context)
	GetMyJobs(context *gin.Context)
	GetSingleJob(context *gin.Context)
}

type jobController struct {
	jobService service.JobService
	logger *logger.Logger
	jwtService service.JwtService
}

func NewJobController(jobService service.JobService, jwtService service.JwtService, logger *logger.Logger) *jobController {
   return &jobController{
	jobService: jobService,
	jwtService: jwtService,
	logger: logger,
   }
}

func (controller *jobController) SaveJob(context *gin.Context) {
	jobDto := dto.Job{}

	userId, _ := context.Get("user_id")

	err := context.ShouldBindJSON(&jobDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))

		return
	}

	_, job := controller.jobService.SaveJob(jobDto, userId.(uint))
	jobSerializer := serializer.JobSerializer{
		Job: job,
	}

	context.JSON(http.StatusCreated, utils.GetResponse(http.StatusCreated, "A new job has been created",jobSerializer.Response() ))
}

func (controller *jobController) GetMyJobs(context *gin.Context) {
	userId, _ := context.Get("user_id")

	_, jobs := controller.jobService.FindJobsByRecruiter(userId.(uint))

	jobsSerializer := serializer.JobsSerializer {
		Jobs: jobs,
	}

	context.JSON(
		http.StatusOK,
		utils.GetResponse(
			http.StatusOK,
			"Jobs retrieved",
			jobsSerializer.Response(),
		),
	)
}

func (controller *jobController) GetSingleJob(context *gin.Context) {
	jobId := context.Param("id")

	jobIdD, _ := strconv.ParseUint(jobId, 10, 12)

	_, job := controller.jobService.GetSingleJob(uint(jobIdD))

	if job.ID == 0 {
		context.JSON(
			http.StatusNotFound,
			utils.GetResponse(http.StatusNotFound, "Job not found", nil),
		)

		return
	}

	jobSerializer := serializer.JobSerializer {
		Job: job,
	}

	context.JSON(
		http.StatusOK,
		utils.GetResponse(
			http.StatusOK,
			"Job found",
			jobSerializer.Response(),
		),
	)

}
