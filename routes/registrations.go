package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"kodski.com/events-api/models"
	"kodski.com/events-api/utils"
)

func createRegistration(c *gin.Context) {
	jwtAuth, isAuthorized := c.Get("jwtAuth")
	if !isAuthorized {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Unauthorized"},
		)
		return
	}

	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	claims := jwtAuth.(*utils.JWTAuth)
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	registration := models.Registration{
		EventID: event.ID,
		UserID: claims.UserId,
	}
	
	err = registration.Save()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		registration,
	)
}

func deleteRegistration(c *gin.Context) {
	jwtAuth, isAuthorized := c.Get("jwtAuth")
	if !isAuthorized {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "Unauthorized"},
		)
		return
	}

	eventId, err := strconv.ParseInt(c.Param("eventId"), 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	claims := jwtAuth.(*utils.JWTAuth)
	err = models.DeleteRegistration(eventId, claims.UserId)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusNoContent,
		nil,
	)
}