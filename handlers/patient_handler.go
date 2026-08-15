package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/backend/config"
	"github.com/yourname/backend/models/request"
	"github.com/yourname/backend/service"
)

type PatientHandler struct {
	service service.PatientService
}

func (h *PatientHandler) SearchPatientId(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		if idParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
			return
		}

		patient, err := h.service.GetPatientByID(idParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if patient == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}

func (h *PatientHandler) SearchPatientList(cfg *config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {
		req := &request.PatientRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("request : ", req)

		patient, err := h.service.FindPatientsList(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if patient == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}
