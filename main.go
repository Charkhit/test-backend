package main

import (
	"net/http"

	"github.com/yourname/backend/config"
	"github.com/yourname/backend/database"
	"github.com/yourname/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	routes.SetupRoutes(r, cfg)

	r.Run(":" + cfg.ServerPort)
}
