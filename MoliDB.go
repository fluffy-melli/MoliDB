package main

import (
	"log"

	"github.com/fluffy-melli/MoliDB/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server = router.SetupRestAPI(server)
	server.Run(":17233")
}
