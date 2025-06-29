package router

import (
	"github.com/Felipe-Janazi/gopportunities/handler"
	"github.com/gin-gonic/gin"
	docs "github.com/Felipe-Janazi/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// Initializer handler
	handler.InitializerHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath= basePath

	// Request group
	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	// Initialize swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
