package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	godotenv.Load("dev.env")
}

func getTestDb(dbUrl string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, err
	}

	//Instantiates the context and connects the to the client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("TEST_DB_NAME"))
	if err != nil {
		return nil, err
	}
	if db == nil {
		return db, fmt.Errorf("db is nil")
	}
	return db, nil
}
func TestGetAllPost(t *testing.T) {
	db, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}

	postRepo := NewPostRepository(db)
	posts, err := postRepo.GetAllPosts()
	if posts == nil || len(posts) == 0 || err != nil {
		t.Errorf("There were no posts returned!")
	}
}

func TestGetExistingPost(t *testing.T) {
	postId := "1"
	db, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postRepo := NewPostRepository(db)
	post, err := postRepo.FindPost(postId)
	if err != nil {
		t.Errorf("There was an error getting the post: %v", err)
	}
	if post.ID != postId {
		t.Errorf("The post returned was not the one requested!")
	}
}

func TestPostDoesNotExist(t *testing.T) {
	postId := "100000"
	db, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postRepo := NewPostRepository(db)
	_, err = postRepo.FindPost(postId)
	if err == nil {
		t.Errorf("There was no error when there should have been!")
	}

}

//test getting a user
func TestGetUser(t *testing.T) {
	userId := "1"
	db, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	userRepo := NewUserRepository(db)
	user, err := userRepo.FindUser(userId)
	if err != nil {
		t.Errorf("There was an error getting the user: %v", err)
	}
	if user.UserId != userId {
		t.Errorf("The user returned was not the one requested!")
	}
}

//test getting a user that does not exist
func TestGetUserThatDoesNotExist(t *testing.T) {
	userId := "100000"
	db, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	userRepo := NewUserRepository(db)
	_, err = userRepo.FindUser(userId)
	if err != mongo.ErrNoDocuments {
		t.Errorf("There was no error when there should have been!")
	}
}
