package router

import (
	"fmt"
	"go-dpm/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	R          *gin.Engine
	ServerAddr *string
	DB         *database.Database
}

func NewRouter(db *database.Database) *Router {
	serverAddr := fmt.Sprint("127.0.0.1:", os.Getenv("server_addr"))

	router := gin.Default()

	return &Router{
		ServerAddr: &serverAddr,
		R:          router,
		DB:         db,
	}
}

func registerRoutes(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, "user test")
		})
	}
}