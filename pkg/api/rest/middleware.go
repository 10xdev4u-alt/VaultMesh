package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the presence of a valid API key in the headers.
func AuthMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key == "" || key != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
