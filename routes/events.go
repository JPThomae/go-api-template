package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)


func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	} else if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "you do not have permission to update this event"})
		return
	}
	
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.Id = event.Id // Keep the same ID
	event.UserID = 1 // Assuming a default user ID for simplicity

	err = updatedEvent.UpdateEventById()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"error": "you do not have permission to update this event"})
		return
	}

	err = event.DeleteEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}

func registerForEvent(context *gin.Context) {

}

func unregisterFromEvent(context *gin.Context) {
	
}
