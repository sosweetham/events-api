package routes

import (
	"github.com/gin-gonic/gin"
	"kodski.com/events-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:eventId", updateEvent)
	server.DELETE("/events/:eventId", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}