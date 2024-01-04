package route

import (
	"github.com/eliemugenzi/simply-hired/controller"
	"github.com/eliemugenzi/simply-hired/middleware"
	repository "github.com/eliemugenzi/simply-hired/repositories"
	service "github.com/eliemugenzi/simply-hired/services"
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JobRoute(db *gorm.DB, jobRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		authRepository repository.AuthRepo = repository.NewAuthRepo(db)
		jwtService service.JwtService = service.NewJwtService()
		jobRepository repository.JobRepo = repository.NewJobRepo(db)
		jobService service.JobService = service.NewJobService(jobRepository, authRepository)
		jobController controller.JobController = controller.NewJobController(jobService, jwtService, logger)
		roleMiddleware middleware.RoleMiddleware = middleware.NewRoleMiddleware(authRepository)
	)

	jobRouter.POST("/", roleMiddleware.AuthorizeRole("HR"), jobController.SaveJob)
	jobRouter.GET("/myjobs", roleMiddleware.AuthorizeRole("HR"), jobController.GetMyJobs)
	jobRouter.GET("/:id", jobController.GetSingleJob)
}