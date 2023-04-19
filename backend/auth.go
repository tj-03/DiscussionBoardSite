package main

import (
	"errors"
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			//c.AbortWithStatusJSON(401, "User does not exist")
			u, err := provider.FirebaseClient.GetUser(c, token.UID)
			if err != nil {
				c.AbortWithStatusJSON(401, "User does not exist")
				return
			}
			email := u.Email
			provider.users.CreateNewUser(User{UserId: token.UID, Username: email})
			return
		}
		c.Next()
	}
}
