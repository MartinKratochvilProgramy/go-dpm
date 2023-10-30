package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) removeUnusedStocks(c *gin.Context) {
	err := r.DB.RemoveUnusedStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unused stocks removed succesfully."})
	return
}
