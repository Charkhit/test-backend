package handlers

import (
	"fmt"
	"net/http"

	"github.com/chakhit/backend/config"
	"github.com/chakhit/backend/models/request"
	"github.com/chakhit/backend/service"
	"github.com/gin-gonic/gin"
)

type PeopleHandler struct {
	service service.PeopleService
}

func NewPeopleHandler(service service.PeopleService) *PeopleHandler {
	return &PeopleHandler{service: service}
}

func (h *PeopleHandler) AddPeople(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {
		req := &request.PeopleAddRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("request: %+v\n", *req)

		err := h.service.AddPeoples(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Peoples added successfully"})
	}
}

func (h *PeopleHandler) GetAllPeople(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {
		data, err := h.service.GetPeopleList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
