package router

import (
	docs "github.com/forbatnew/gopportunities/docs"
	"github.com/forbatnew/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Title = "Project Gopportunities"
	docs.SwaggerInfo.Description = "This is a sample opening server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	v1 := router.Group(basePath)
	{
		//Show Opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	//Initialize Swagger

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
