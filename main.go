package main

import (
	"os"

	"github.com/eliemugenzi/simply-hired/db/config"
	"github.com/eliemugenzi/simply-hired/middleware"
	"github.com/eliemugenzi/simply-hired/route"
	"github.com/eliemugenzi/simply-hired/utils"
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)


type Repository struct {
	DB *gorm.DB
}

var db *gorm.DB = config.Configure()


func main() {

	godotenv.Load()
	defer config.CloseConnection(db)
	logger := logger.NewLogger()

	router := gin.Default()

	zLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	router.Use(middleware.RequestLogger(&zLogger))

	route.RootRoute(db, router, logger)

	utils.InitializeCustomValidations()

	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H {
			"message": "Hello world",
		})
	})
	router.Run(":"+ os.Getenv("APP_PORT"))
}