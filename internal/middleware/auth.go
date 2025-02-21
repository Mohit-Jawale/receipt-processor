package middleware

import (
	"net/http"
	"receipt-processor/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware ensures that the request contains a valid JWT token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store the username in the request context
		c.Set("username", claims.Username)
		c.Next()
	}
}
