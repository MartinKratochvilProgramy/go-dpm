package router

import (
	"go-dpm/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Run() {
	r.R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "message")
	})

	user := r.R.Group("/user")
	{
		user.POST("/register", r.register)
		user.POST("/login", r.login)
	}

	portfolio := r.R.Group("/portfolio")
	portfolio.Use(middleware.JwtAuthMiddleware())
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
