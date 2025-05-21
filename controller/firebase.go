package controller

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// Verify ID token
func verifyIDToken(app *firebase.App, idToken string) (*auth.Token, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Middleware to handle CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Example route to store user data
func StoreUserData(app *firebase.App, c *gin.Context) {
	// Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	idToken := strings.TrimPrefix(authHeader, "Bearer ")

	// Verify the token
	token, err := verifyIDToken(app, idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to verify ID token"})
		return
	}

	email, ok := token.Claims["email"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found in token"})
		return
	}

	log.Printf("user email: %s", email)

	c.JSON(http.StatusOK, gin.H{"message": "User data stored successfully!"})
}
