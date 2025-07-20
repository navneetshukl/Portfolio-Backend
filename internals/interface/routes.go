package routes

import (
	"portfolio/internals/interface/api/rest/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(eh *handlers.Handler) *gin.Engine {
	// Initialize a new Gin router
	router := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "https://navneetdev.netlify.app/"} // Add your frontend URLs
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.MaxAge = 12 * 60 * 60 // Cache CORS preflight for 12 hours

	// Apply CORS middleware to the router
	router.Use(cors.New(config))

	mux := router.Group("/api")
	mux.POST("/sendEmail", eh.SendEmailHandler())
	return router
}
