package router

import (
	docs "github.com/fabioaacarneiro/gopportunities/docs"
	"github.com/fabioaacarneiro/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
    // initialize handler
    handler.InitializeHandler()

    basePath := "/api/v1"
    docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

    // initilize Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
