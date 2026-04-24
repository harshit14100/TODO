package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshit14100/go-todo/database/dbHelper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("Authorization")
		if sessionID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session."})
			return
		}

		userID, err := dbHelper.GetUserBySession(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session."})
			return
		}

		c.Set("userID", userID) // store the userID in the context so handlers can grab it
		c.Next()                // Let the req continue to the header
	}
}
