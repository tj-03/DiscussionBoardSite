package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("dev.env")
}

func TestGetAllPost(t *testing.T) {
	db, err := getDbConn(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db is nil")
	}
	postRepo := NewPostRepository(db)
	posts, err := postRepo.GetAllPosts()
	if posts == nil || len(posts) == 0 {
		t.Errorf("There were no posts returned!")
	}

}
