package controller

import (
	"net/http"

	"github.com/eliemugenzi/simply-hired/dto"
	service "github.com/eliemugenzi/simply-hired/services"
	"github.com/eliemugenzi/simply-hired/utils"
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
)

type JobController interface {
	SaveJob(context *gin.Context)
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

	context.JSON(http.StatusCreated, utils.GetResponse(http.StatusCreated, "A new job has been created",job ))
}