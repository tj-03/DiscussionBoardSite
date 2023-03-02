package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func checkIfDBExists(client *mongo.Client, dbName string) bool {
	dbs, err := client.ListDatabaseNames(context.Background(), map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	for _, db := range dbs {
		if db == dbName {
			return true
		}
	}
	return false
}
func createDatabaseCollections(myDB *mongo.Database) {
	postsCol := myDB.Collection("Posts")
	myDB.Collection("Comments")
	usersCol := myDB.Collection("Users")

	mockUsers := createTestUsers()
	for _, user := range mockUsers {
		_, err := usersCol.InsertOne(context.Background(), user)
		if err != nil {
			log.Fatal(err)
		}
	}
	mockPosts := createTestPosts()
	for _, post := range mockPosts {
		_, err := postsCol.InsertOne(context.Background(), post)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func createTestDatabase(client *mongo.Client) {
	dbs, err := client.ListDatabaseNames(context.Background(), map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
	for _, db := range dbs {
		if db == "TEST_DB" {
			return
		}
	}
	if checkIfDBExists(client, "TEST_DB") {
		return
	}
	myDB := client.Database("TEST_DB")
	createDatabaseCollections(myDB)

}
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
	createTestDatabase(client)
	//Creates database and collections
	if checkIfDBExists(client, "myDB") {
		return
	}
	myDB := client.Database("myDB")
	myDB.CreateCollection(context.Background(), "Posts")
	myDB.CreateCollection(context.Background(), "Comments")
	myDB.CreateCollection(context.Background(), "Users")

	if err != nil {
		log.Fatal(err)
	}

}
