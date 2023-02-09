package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	//Instantiates the client and connection location
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
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

	//Disconnects the client
	defer client.Disconnect(ctx)

	//Creates database and collections
	myDB := client.Database("myDB")
	postCollection := myDB.Collection("Posts")
	commentCollection := myDB.Collection("Comments")
	userCollection := myDB.Collection("Users")

	//Creates test post in post collection
	_, err = postCollection.InsertOne(ctx, bson.D{
		{Key: "Title", Value: "Test Title"},
		{Key: "Author", Value: "Test Author"},
		{Key: "Body", Value: "This is a test post."},
		{Key: "PostID", Value: "0"},
	})
	if err != nil {
		log.Fatal(err)
	}

	//Creates test comment in comment collection
	_, err = commentCollection.InsertOne(ctx, bson.D{
		{Key: "PostID", Value: "0"},
		{Key: "Author", Value: "Test Author"},
		{Key: "Body", Value: "This is a test comment on the test post."},
		{Key: "CommentID", Value: "0"},
	})
	if err != nil {
		log.Fatal(err)
	}

	//Create user database
	_, err = userCollection.InsertOne(ctx, bson.D{
		{Key: "UserID", Value: ""},
		{Key: "Username", Value: "TestAccount"},
	})
	if err != nil {
		log.Fatal(err)
	}

	//Lists the current databases
	DBs, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DBs)

}
