package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("admin_username"): os.Getenv("admin_password"),
	})
}
