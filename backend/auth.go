package main

import (
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type FirebaseAuthProvider struct {
	FirebaseClient *auth.Client
	users          UserRepository
}

func NewFirebaseAuthProvider(firebaseClient *auth.Client, db *mongo.Database) FirebaseAuthProvider {
	return FirebaseAuthProvider{
		FirebaseClient: firebaseClient,
		users:          NewUserRepository(db)}
}

func (provider *FirebaseAuthProvider) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authenticating")
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(401, "Invalid auth token")
			return
		}
		token, err := provider.FirebaseClient.VerifyIDToken(c, authHeader[0])
		if err != nil {
			c.AbortWithStatusJSON(401, "Unauthorized")
			return
		}

		_, err = provider.users.FindUser(token.UID)
		if err == mongo.ErrNilDocument {
			c.AbortWithStatusJSON(401, "User does not exist")
			//create user
			return
		}
		c.Next()
	}
}
