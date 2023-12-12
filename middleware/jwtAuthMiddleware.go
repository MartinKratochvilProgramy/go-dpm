package middleware

import (
	"go-dpm/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	// check if encoded username in token and username provided in headers match
	return func(c *gin.Context) {
		err := token.ValidateToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		usernameFromToken, err := token.ExtractTokenUsername(c)
		if usernameFromToken == "" {
			c.String(http.StatusBadRequest, "Username missing in Token.")
			c.Abort()
			return
		}
		usernameFromHeader := c.Request.Header.Get("username")
		if usernameFromHeader == "" {
			c.String(http.StatusBadRequest, "Username missing in Header.")
			c.Abort()
			return
		}

		if usernameFromToken != usernameFromHeader {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
