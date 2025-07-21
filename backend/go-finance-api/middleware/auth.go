package middleware

import (
	"net/http"
	"strings"

	"github.com/davidwang/go-finance-api/go-finance-api/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates a Gin middleware for JWT authentication
func AuthMiddleware(config auth.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		// Check for Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be in format: Bearer {token}"})
			return
		}

		// Extract the token
		tokenString := parts[1]

		// Validate the token
		claims, err := auth.ValidateToken(tokenString, config)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Set the user ID in the context for later use
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)

		c.Next()
	}
}

// Optional middleware that checks for token but allows request to proceed if no token
func OptionalAuthMiddleware(config auth.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// No auth header, continue without setting user info
			c.Next()
			return
		}

		// Check for Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Invalid format, continue without setting user info
			c.Next()
			return
		}

		// Extract the token
		tokenString := parts[1]

		// Validate the token
		claims, err := auth.ValidateToken(tokenString, config)
		if err == nil {
			// Valid token, set user info
			c.Set("userID", claims.UserID)
			c.Set("userEmail", claims.Email)
		}

		c.Next()
	}
}
