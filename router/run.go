package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Run() {
	r.R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "message")
	})

	r.R.POST("/foo", r.foo)

	r.R.POST("/register", r.register)

	r.R.Run(*r.ServerAddr)
}
