package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var NoFilter = map[string]interface{}{}

type User struct {
	UserId   string `json:"id" bson:"_id,omitempty"` // bson is for mongo
	Username string `json:"name" bson:"Username,omitempty"`
}

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{coll: db.Collection("Users")}
}

func (u *UserRepository) FindUser(id string) (User, error) {
	user := User{}
	err := u.coll.FindOne(context.Background(), User{UserId: id}).Decode(&user)
	return user, err
}

func (u *UserRepository) CreateNewUser(user User) error {
	_, err := u.coll.InsertOne(context.Background(), user)
	return err
}

type Post struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	AuthorId string `json:"author_id" bson:"author_id,omitempty"`
	Content  string `json:"content" bson:"content,omitempty"`
}

type PostRepository struct {
	coll *mongo.Collection
}

func NewPostRepository(db *mongo.Database) PostRepository {
	return PostRepository{coll: db.Collection("Posts")}
}

func (p *PostRepository) AddPost(post Post) error {
	_, err := p.coll.InsertOne(context.Background(), post)
	return err
}

func (p *PostRepository) FindPost(postId string) (Post, error) {
	post := Post{}
	err := p.coll.FindOne(context.Background(), bson.M{"_id": postId}).Decode(&post)
	return post, err
}

func (p *PostRepository) GetAllPosts() ([]Post, error) {
	cursor, err := p.coll.Find(context.Background(), NoFilter)
	if err != nil {
		return nil, err
	}
	var posts []Post
	err = cursor.All(context.Background(), &posts)
	return posts, err
}

func (p *PostRepository) GetPostsFromUserId(userId string) ([]Post, error) {
	cursor, err := p.coll.Find(context.Background(), Post{AuthorId: userId})
	if err != nil {
		return nil, err
	}
	var posts []Post
	err = cursor.All(context.Background(), &posts)
	return posts, err
}

type Comment struct {
	PostID    string `json:"post_id" bson:"PostID,omitempty"`
	Author    string `json:"author" bson:"Author,omitempty"`
	Body      string `json:"content" bson:"Body,omitempty"`
	CommentID string `json:"comment_id" bson:"CommentID,omitempty"`
}

type CommentRepository struct {
	coll *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) CommentRepository {
	return CommentRepository{coll: db.Collection("Comments")}
}

func (c *CommentRepository) AddCommet(comment Comment) error {
	_, err := c.coll.InsertOne(context.Background(), comment)
	return err
}

func (c *CommentRepository) GetAllComments() ([]Comment, error) {
	cursor, err := c.coll.Find(context.Background(), NoFilter)
	if err != nil {
		return nil, err
	}
	var comments []Comment
	err = cursor.All(context.Background(), &comments)
	return comments, err
}
