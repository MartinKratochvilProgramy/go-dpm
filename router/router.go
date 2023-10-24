package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct {
	R          *gin.Engine
	ServerAddr *string
}

func NewRouter() *Router {
	serverAddr := fmt.Sprint(":", os.Getenv("server_addr"))

	router := gin.Default()

	registerRoutes(router)

	return &Router{
		R:          router,
		ServerAddr: &serverAddr,
	}
}

func registerRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Millisecond)
		c.JSON(http.StatusOK, "message")
	})
}
