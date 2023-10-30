package router

import (
	"go-dpm/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Run() {
	r.R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	user := r.R.Group("/user")
	{
		user.POST("/register", r.register)
		user.POST("/login", r.login)
	}

	portfolio := r.R.Group("/portfolio")
	portfolio.Use(middleware.JwtAuthMiddleware())
	{
		portfolio.GET("/get", r.getPortfolio)
		portfolio.POST("/add", r.addStockToPortfolio)
		portfolio.POST("/remove", r.removeStockFromPortfolio)
	}

	stocks := r.R.Group("/stocks", r.adminAuth)
	{
		stocks.PUT("/update", r.updateStocks)
		stocks.PUT("/remove_unused", r.removeUnusedStocks)
	}

	conversionRates := r.R.Group("/conversion_rates", r.adminAuth)
	{
		conversionRates.PUT("/update", r.updateConversionRates)
	}

	r.R.Run(*r.ServerAddr)
}
