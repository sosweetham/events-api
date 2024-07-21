package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"kodski.com/events-api/models"
	"kodski.com/events-api/utils"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(
		http.StatusCreated,
		user,
	)
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	err = user.Authenticate()
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	// token, err := utils.GenerateToken(user.Email, user.ID)
	token, err := utils.
		NewJWTAuth(user.Email, user.ID, time.Now().Add(time.Hour * 2).Unix()).
		GenerateToken()
	

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "Logged in", "token": token},
	)
}