package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	//Event Routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	//User ROutes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
