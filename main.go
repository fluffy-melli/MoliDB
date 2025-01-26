package main

import (
	"time"

	"github.com/fluffy-melli/MoliDB/internal/middleware"
	"github.com/fluffy-melli/MoliDB/internal/router"
	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/fluffy-melli/MoliDB/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/fluffy-melli/MoliDB/docs"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		logger.ERROR("Error loading .env file: %v", err)
	}
	config := middleware.MiddlewareConfig{
		Capacity:   750,
		RefillRate: 500.0 / 60.0,
	}
	server := gin.Default()
	server = router.SetupRestAPI(server, config)
	go func() {
		logger.INFO("MoliDB HTTP Starting...")
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
