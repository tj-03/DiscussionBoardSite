package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	godotenv.Load("dev.env")
}

func getTestDb(dbUrl string) (*mongo.Database, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, nil, err
	}

	//Instantiates the context and connects the to the client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, ctx, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, ctx, err
	}

	db := client.Database(os.Getenv("TEST_DB_NAME"))
	if err != nil {
		return nil, ctx, err
	}
	if db == nil {
		return db, ctx, fmt.Errorf("db is nil")
	}
	return db, ctx, nil
}
func TestGetAllPosts(t *testing.T) {
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}

	postRepo := NewPostRepository(db)
	posts, err := postRepo.GetAllPosts()
	if posts == nil || len(posts) != 5 || err != nil {
		t.Errorf("Incorrect number of posts returned!")
	}
}

func TestGetExistingPost(t *testing.T) {
	postId := "1"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
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
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postRepo := NewPostRepository(db)
	_, err = postRepo.FindPost(postId)
	if err == nil {
		t.Errorf("There was no error when there should have been!")
	}

}

func TestAddPost(t *testing.T) {
	postId := "123456789abcdef"
	author := "2"
	body := "Add Post Test"

	var nPost Post
	nPost.AuthorId = author
	nPost.Content = body
	nPost.ID = postId

	db, ctx, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postRepo := NewPostRepository(db)
	_, err = postRepo.FindPost(postId)
	if err == nil {
		t.Errorf("Post should not exist!")
	}

	postRepo.AddPost(nPost)

	post, err := postRepo.FindPost(postId)
	if err != nil {
		t.Errorf("There was an error getting the post: %v", err)
	}
	if post.ID != postId {
		t.Errorf("The post returned was not the one requested!")
	}

	postRepo.coll.DeleteOne(ctx, bson.D{{"_id", "123456789abcdef"}})
}

func TestGetPostFromUser(t *testing.T) {
	postId := "1"
	userId := "1"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postRepo := NewPostRepository(db)
	post, err := postRepo.FindPostFromUserID(postId, userId)
	if err != nil {
		t.Errorf("There was an error getting the post: %v", err)
	}
	if post.ID != postId || post.AuthorId != userId {
		t.Errorf("The post returned was not the one requested!")
	}
}

func TestGetPostFromTitle(t *testing.T) {
	title := "Title2"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postRepo := NewPostRepository(db)
	post, err := postRepo.FindPostFromTitle(title)
	if err != nil {
		t.Errorf("There was an error getting the post: %v", err)
	}
	if post.Title != title {
		t.Errorf("The post returned was not the one requested!")
	}
}

func TestGetAllPostFromUser(t *testing.T) {
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	userId := "1"

	postRepo := NewPostRepository(db)
	posts, err := postRepo.GetAllPostsFromUserId(userId)
	if posts == nil || len(posts) != 2 || err != nil {
		t.Errorf("Incorrect number of posts returned!")
	}
}
func TestGetAllUser(t *testing.T) {
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}

	userRepo := NewUserRepository(db)
	users, err := userRepo.GetAllUsers()
	if users == nil || len(users) != 5 || err != nil {
		t.Errorf("Incorrec number of users returned!")
	}
}

// test getting a user
func TestGetUser(t *testing.T) {
	userId := "1"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
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

// test getting a user that does not exist
func TestGetUserThatDoesNotExist(t *testing.T) {
	userId := "100000"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	userRepo := NewUserRepository(db)
	_, err = userRepo.FindUser(userId)
	if err != mongo.ErrNoDocuments {
		t.Errorf("There was no error when there should have been!")
	}
}

func TestAddUser(t *testing.T) {
	userId := "123456789abcdef"
	name := "TestAddUser"

	var nUser User
	nUser.UserId = userId
	nUser.Username = name

	db, ctx, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	userRepo := NewUserRepository(db)
	_, err = userRepo.FindUser(userId)
	if err == nil {
		t.Errorf("User should not exist!")
	}

	userRepo.CreateNewUser(nUser)

	user, err := userRepo.FindUser(userId)
	if err != nil {
		t.Errorf("There was an error getting the user: %v", err)
	}
	if user.UserId != userId {
		t.Errorf("The user returned was not the one requested!")
	}

	userRepo.coll.DeleteOne(ctx, bson.D{{"_id", userId}})
}

func TestGetAllComments(t *testing.T) {
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}

	commentRepo := NewCommentRepository(db)
	comments, err := commentRepo.GetAllComments()
	if comments == nil || len(comments) != 3 || err != nil {
		t.Errorf("Incorrect number of comments returned!")
	}
}

func TestGetAllCommentsFromUser(t *testing.T) {
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	userId := "1"

	commentRepo := NewCommentRepository(db)
	comments, err := commentRepo.GetAllCommentsFromUserId(userId)
	if comments == nil || len(comments) != 2 || err != nil {
		t.Errorf("Incorrect number of comments returned!")
	}
}

func TestGetAllCommentsFromPost(t *testing.T) {
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	postId := "2"

	commentRepo := NewCommentRepository(db)
	comments, err := commentRepo.GetAllCommentsFromPostId(postId)
	if comments == nil || len(comments) != 1 || err != nil {
		t.Errorf("Incorrect number of comments returned!")
	}
}

func TestGetExistingComment(t *testing.T) {
	postId := "1"
	commentId := "1"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	commentRepo := NewCommentRepository(db)
	comment, err := commentRepo.FindComment(postId, commentId)
	if err != nil {
		t.Errorf("There was an error getting the post: %v", err)
	}
	if comment.CommentID != commentId {
		t.Errorf("The comment returned was not the one requested!")
	}
}

func TestCommentDoesNotExist(t *testing.T) {
	postId := "1"
	commentId := "3"
	db, _, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	commentRepo := NewCommentRepository(db)
	_, err = commentRepo.FindComment(postId, commentId)
	if err == nil {
		t.Errorf("There was no error when there should have been!")
	}

}

func TestAddComment(t *testing.T) {
	postId := "3"
	author := "4"
	body := "Test Add Comment"
	commentId := "123456789abcdef"

	var nComment Comment
	nComment.PostID = postId
	nComment.Author = author
	nComment.Body = body
	nComment.CommentID = commentId

	db, ctx, err := getTestDb(os.Getenv("DB_URL"))
	if err != nil {
		t.Errorf("There was an error connecting to the database: %v", err)
	}
	commentRepo := NewCommentRepository(db)
	_, err = commentRepo.FindComment(postId, commentId)
	if err == nil {
		t.Errorf("Comment should not exist!")
	}

	commentRepo.AddComment(nComment)

	comment, err := commentRepo.FindComment(postId, commentId)
	if err != nil {
		t.Errorf("There was an error getting the comment: %v", err)
	}
	if comment.CommentID != commentId {
		t.Errorf("The comment returned was not the one requested!")
	}

	commentRepo.coll.DeleteOne(ctx, bson.D{{"PostID", "3"}, {"CommentID", commentId}})
}
