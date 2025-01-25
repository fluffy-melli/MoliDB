package router

import (
	"github.com/fluffy-melli/MoliDB/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRestAPI(router *gin.Engine) *gin.Engine {
	router.GET("/collection", middleware.Collection)
	router.GET("/collection/:id", middleware.CollectionID)
	router.PUT("/collection/:id", middleware.CollectionPut)
	router.DELETE("/collection/:id", middleware.CollectionDelete)
	return router
}
