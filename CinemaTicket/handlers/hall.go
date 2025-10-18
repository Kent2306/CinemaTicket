package handlers

import (
	"cinematicket/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHalls(c *gin.Context) {
	halls, err := models.GetAllHalls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, halls)
}
