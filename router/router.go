package router

import (
	"fmt"
	"go-dpm/database"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	R          *gin.Engine
	ServerAddr *string
	DB         *database.Database
	AdminAuth  gin.HandlerFunc
}

func NewRouter(db *database.Database) *Router {
	serverAddr := fmt.Sprint("127.0.0.1:", os.Getenv("server_addr"))
	AdminAuth := gin.BasicAuth(gin.Accounts{
		os.Getenv("admin_username"): os.Getenv("admin_password"),
	})

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	return &Router{
		ServerAddr: &serverAddr,
		R:          router,
		DB:         db,
		AdminAuth:  AdminAuth,
	}
}
