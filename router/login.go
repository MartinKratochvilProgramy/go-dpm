package router

import (
	"go-dpm/utils/bcrypt"
	"go-dpm/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) login(c *gin.Context) {
	var body struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.DB.GetUser(body.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found."})
		return
	}

	if !bcrypt.ComparePasswords(body.Password, user.PasswordHash) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password."})
		return
	}

	token, err := token.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User login successful.", "token": token})
}
