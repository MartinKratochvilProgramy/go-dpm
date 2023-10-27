package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Run() {
	r.R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "message")
	})

	r.R.POST("/register", r.register)

	v1 := r.R.Group("/portfolio")
	{
		v1.POST("/get_portfolio", r.getPortfolio)
		v1.POST("/add_stock", r.addStockToPortfolio)
		v1.POST("/remove_stock", r.removeStockFromPortfolio)
	}

	r.R.Run(*r.ServerAddr)
}
