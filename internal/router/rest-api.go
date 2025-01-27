package router

import (
	"github.com/fluffy-melli/MoliDB/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRestAPI(router *gin.Engine, config middleware.MiddlewareConfig) *gin.Engine {
	router.Use(middleware.AllowIPMiddleware())
	router.Use(middleware.TokenBucketMiddleware(config))
	router.GET("/collection", middleware.Collection)
	router.GET("/collection/:id", middleware.CollectionID)
	router.PUT("/collection/:id", middleware.CollectionPut)
	router.DELETE("/collection/:id", middleware.CollectionDelete)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
