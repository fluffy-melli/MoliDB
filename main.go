package main

import (
	"log"
	"time"

	"github.com/fluffy-melli/MoliDB/internal/router"
	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/fluffy-melli/MoliDB/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server = router.SetupRestAPI(server)
	server.Run(":17233")
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			log.Println("Performing backup...")
			runtime.DB.Backup()
		}
	}()
}
