package routes

import (
	"portfolio/internals/interface/api/rest/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(eh *handlers.Handler) *gin.Engine {
	// Initialize a new Gin router
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://navneetdev.netlify.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}

	// Apply CORS middleware to the router
	router.Use(cors.New(config))

	mux := router.Group("/api")
	mux.POST("/sendEmail", eh.SendEmailHandler())
	return router
}
