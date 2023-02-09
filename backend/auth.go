package main

import (
	"fmt"
	"log"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type FirebaseAuthProvider struct {
	FirebaseClient *auth.Client
}

func NewFirebaseAuthProvider(firebaseClient *auth.Client) FirebaseAuthProvider {
	return FirebaseAuthProvider{FirebaseClient: firebaseClient}
}


//TODO: Change <UserIDType> to the UserID type
func checkUserExists(UserID <UserIDType>) bool {
	//Instantiates the client and connection location
	client, err := mongo.NewClient(options.Client().ApplyURI(server))
	if err != nil {
		log.Fatal(err)
	}

	//Instantiates the context and connects the to the client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	if client.Database("myDB").Collection("Users").Find(ctx, bson.M{"UserID": UserID}){
		return true
	}
	else{
		return false
	}
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
		//See comments for checkUserExists()
		//if checkUserExists(<UserID>){
		//	authenticate users
		//}
		//else{
		//	user does not exist	
		//}

		c.Next()
	}
}
