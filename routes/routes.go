package routes

import (
	"github.com/GitNinja36/backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/health", controllers.HealthCheck)
		api.POST("/favorites", controllers.AddFavorite)
		api.GET("/favorites/:userId", controllers.GetFavorites)
		api.DELETE("/favorites/:userId/:recipeId", controllers.DeleteFavorite)
	}
}
