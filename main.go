package main

import (
	"time"

	"github.com/fluffy-melli/MoliDB/internal/router"
	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/fluffy-melli/MoliDB/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/fluffy-melli/MoliDB/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.ERROR("Error loading .env file: %v", err)
	}
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server = router.SetupRestAPI(server)
	go func() {
		logger.INFO("MoliDB Starting... / Port: 17233")
		if err := server.Run(":17233"); err != nil {
			logger.ERROR("MoliDB Starting Error: %v", err)
		}
	}()
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			logger.INFO("Starting Backup Process...")
			err := runtime.DB.Backup()
			if err != nil {
				logger.WARNING("There was an issue during the backup: %v", err)
			} else {
				logger.INFO("The backup was successful.")
			}
		}
	}()
	select {}
}
