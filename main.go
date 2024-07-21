package main

import (
	"github.com/gin-gonic/gin"
	"kodski.com/events-api/db"
	"kodski.com/events-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}