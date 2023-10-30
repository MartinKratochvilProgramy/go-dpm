package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) getPortfolio(c *gin.Context) {
	username := c.Request.Header.Get("username")
	if username == "" {
		c.String(http.StatusBadRequest, "Username missing in Header.")
		c.Abort()
		return
	}

	pf, err := r.DB.GetPortfolio(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"portfolio": pf})
}
