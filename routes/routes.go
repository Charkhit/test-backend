package routes

import (
	"net/http"
	"strings"

	"github.com/yourname/backend/config"
	"github.com/yourname/backend/handlers"
	"github.com/yourname/backend/utils"

	"github.com/gin-gonic/gin"
)

func AuthRequired(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cfg == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "config is missing"})
			c.Abort()
			return
		}

		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		if _, err := utils.ValidateToken(parts[1], cfg.JWTSecret); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func SetupRoutes(engine *gin.Engine, cfg *config.Config) {
	auth := engine.Group("/staff")
	{
		auth.POST("/create", handlers.Register(cfg))
		auth.POST("/login", handlers.Login(cfg))
	}

	patientHandler := handlers.PatientHandler{}
	patient := engine.Group("/patient")
	{
		patient.Use(AuthRequired(cfg))
		patient.GET("/search/:id", patientHandler.SearchPatientId(cfg))
		patient.POST("/search", patientHandler.SearchPatientList(cfg))
	}
}
