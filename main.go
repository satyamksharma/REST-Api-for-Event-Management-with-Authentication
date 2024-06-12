package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satyamksharma/REST-Api-for-Event-Management-with-Authentication.git/db"
	"github.com/satyamksharma/REST-Api-for-Event-Management-with-Authentication.git/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events",createEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, Please try later!"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the Event, Please try later"})
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the Event, Please try later"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, Please try later!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message":"Event Created!", "event": event})
}

