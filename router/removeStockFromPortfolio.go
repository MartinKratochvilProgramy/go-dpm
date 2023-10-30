package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) removeStockFromPortfolio(c *gin.Context) {
	username := c.Request.Header.Get("username")
	if username == "" {
		c.String(http.StatusBadRequest, "Username missing in Header.")
		c.Abort()
		return
	}

	var body struct {
		Ticker string `json:"ticker" binding:"required"`
		Shares int    `json:"shares" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Shares <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount of shares has to be positive."})
		return
	}

	err := r.DB.RemoveStockFromPortfolio(username, body.Ticker, body.Shares)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf("Succesfully removed %d %s.", body.Shares, body.Ticker)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
