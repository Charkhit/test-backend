package routes

import (
	"github.com/chakhit/backend/config"
	"github.com/chakhit/backend/database"
	"github.com/chakhit/backend/handlers"
	"github.com/chakhit/backend/repository"
	"github.com/chakhit/backend/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine, cfg *config.Config) {

	db := database.Connect(cfg)

	peopleRepo := repository.NewPeopleRepository(db)
	peopleService := service.NewPeopleService(peopleRepo)
	peopleHandler := handlers.NewPeopleHandler(*peopleService)
	people := engine.Group("/people")
	{
		people.POST("/add", peopleHandler.AddPeople(cfg))
		people.POST("/list", peopleHandler.GetAllPeople(cfg))
	}

}
