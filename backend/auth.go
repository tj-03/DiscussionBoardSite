package main

import (
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type FirebaseAuthProvider struct {
	FirebaseClient *auth.Client
}

func NewFirebaseAuthProvider(firebaseClient *auth.Client) FirebaseAuthProvider {
	return FirebaseAuthProvider{FirebaseClient: firebaseClient}
}

func (ap *FirebaseAuthProvider) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authenticating")
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(401, "Invalid auth token")
			return
		}
		_, err := ap.FirebaseClient.VerifyIDToken(c, authHeader[0])
		if err != nil {
			c.AbortWithStatusJSON(401, "Unauthorized")
			return
		}

		//TODO: Check if user exists in DB

		c.Next()
	}
}
