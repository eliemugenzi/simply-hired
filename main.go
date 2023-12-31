package main

import (
	"os"

	"github.com/eliemugenzi/simply-hired/db/config"
	"github.com/eliemugenzi/simply-hired/route"
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type Repository struct {
	DB *gorm.DB
}

var db *gorm.DB = config.Configure()


func main() {
	defer config.CloseConnection(db)
	logger := logger.NewLogger()

	router := gin.Default()

	route.RootRoute(db, router, logger)

	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H {
			"message": "Hello world",
		})
	})
	router.Run(":"+ os.Getenv("APP_PORT"))
}