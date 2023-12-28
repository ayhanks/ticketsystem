package routes

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/api/v1/auth")

	{
		auth.POST("login", HandleLogin)
	}


	return router
}