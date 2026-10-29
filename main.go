package main

import (
	"net/http"

	"github.com/chakhit/backend/config"
	"github.com/chakhit/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	routes.SetupRoutes(r, cfg)

	r.Run(":" + cfg.ServerPort)
}
