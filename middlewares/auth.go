package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kodski.com/events-api/utils"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "Unauthorized"},
		)
		return
	}

	jwtAuth, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "Unauthorized"},
		)
		return
	}

	c.Set("jwtAuth", jwtAuth)
	c.Next()
}