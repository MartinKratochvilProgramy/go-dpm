package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) getPortfolio(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pf, err := r.DB.GetPortfolio(body.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"portfolio": pf})
}
