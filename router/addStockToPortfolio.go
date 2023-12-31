package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) addStockToPortfolio(c *gin.Context) {
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

	err := r.DB.AddStockToPortfolio(username, body.Ticker, body.Shares)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if body.Shares <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount of shares has to be positive."})
		return
	}

	message := fmt.Sprintf("Succesfully added %s %d", body.Ticker, body.Shares)
	c.JSON(http.StatusOK, gin.H{"message": message})
}
