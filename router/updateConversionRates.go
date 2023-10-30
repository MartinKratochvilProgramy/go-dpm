package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) updateConversionRates(c *gin.Context) {
	err := r.DB.UpdateConversionRates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversion rates updated succesfully."})
	return
}
