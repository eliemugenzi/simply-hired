package route

import (
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(db *gorm.DB, router *gin.Engine, logger *logger.Logger) {
   apiRouter := router.Group("/api/v1")
   
   // Auth router config
   authRouter := apiRouter.Group("/auth")
   AuthRoute(db, authRouter, logger)

   // Job router config
   jobRouter := apiRouter.Group("/jobs")
   JobRoute(db, jobRouter, logger)
}
