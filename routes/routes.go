package routes

import (
	"github.com/gin-gonic/gin"
	"kodski.com/events-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)
	authenticated.POST("/events/:id/register", createRegistration)
	authenticated.DELETE("/events/:id/register", deleteRegistration)
	
	server.POST("/signup", signup)
	server.POST("/login", login)
}