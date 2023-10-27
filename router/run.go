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

	portfolio := r.R.Group("/portfolio")
	{
		portfolio.POST("/get", r.getPortfolio)
		portfolio.POST("/add", r.addStockToPortfolio)
		portfolio.POST("/remove", r.removeStockFromPortfolio)
	}

	stocks := r.R.Group("/stocks")
	{
		stocks.PUT("/update", r.updateStocks)
	}

	r.R.Run(*r.ServerAddr)
}
