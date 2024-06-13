package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/satyamksharma/REST-Api-for-Event-Management-with-Authentication.git/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	//Event Routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/", middleware.Authenticate)
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//User ROutes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
