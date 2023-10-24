package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) foo(c *gin.Context) {
	var body struct {
		Foo string `json:"foo"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprint("Request received successfully: ", body.Foo),
	})
}
