package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Currency string `json:"currency"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.DB.CreateUser(body.Username, body.Password, body.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created succesfully!"})
}
