package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	godotenv.Load("dev.env")
}
func servConn(server Server) (*mongo.Client, context.Context) {
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

	return client, ctx
}
func main() {
	server := NewServer(os.Getenv("DB_URL"))
	//log to a file
	log.Fatal(server.Run(os.Getenv("PORT")))

	//Instantiates the client and connection location
	client, ctx := servConn(server)
	//Disconnects the client
	defer client.Disconnect(ctx)
}
