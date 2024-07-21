package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"kodski.com/events-api/models"
	"kodski.com/events-api/utils"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		events,
	)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		event,
	)
}

func createEvent(c *gin.Context) {
	var event models.Event
	jwtAuth, isAuthorized := c.Get("jwtAuth")
	if !isAuthorized {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Unauthorized"},
		)
		return
	}
	claims := jwtAuth.(*utils.JWTAuth)
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	event.UserID = claims.UserId
	err = event.Save()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		gin.H{"event": event},
	)
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	updatedEvent.ID = eventId
	updatedEvent.UpdateEvent()
	c.JSON(
		http.StatusOK,
		gin.H{"event": updatedEvent},
	)
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{"message": "Event deleted successfully"},
	)
}