package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) updateStocks(c *gin.Context) {
	err := r.DB.UpdateStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stocks updated succesfully."})
}
