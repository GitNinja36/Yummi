package main

import (
	"os"

	"github.com/GitNinja36/backend/config"
	"github.com/GitNinja36/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	if os.Getenv("NODE_ENV") == "production" {
		config.StartCronJob()
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router.Run(":" + port)
}
