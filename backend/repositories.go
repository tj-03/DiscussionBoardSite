package main

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u *UserRepository) GetAllUsers() ([]User, error) {
	cursor, err := u.coll.Find(context.Background(), NoFilter)
	if err != nil {
		return nil, err
	}
	users := []User{}
	err = cursor.All(context.Background(), &users)
	return users, err
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

func (u *UserRepository) FindUsername(username string) (User, error) {
	user := User{}
	err := u.coll.FindOne(context.Background(), User{Username: username}).Decode(&user)
	return user, err
}

type Post struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title,omitempty"`
	AuthorId string `json:"author_id" bson:"author_id,omitempty"`
	Content  string `json:"content" bson:"content,omitempty"`
	Points   int    `json:"points" bson:"points,omitempty"`
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
	objId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return post, err
	}
	err = p.coll.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&post)
	return post, err
}

func (p *PostRepository) GetAllPosts() ([]Post, error) {
	cursor, err := p.coll.Find(context.Background(), NoFilter)
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	err = cursor.All(context.Background(), &posts)
	return posts, err
}

func (p *PostRepository) GetAllPostsFromUserId(userId string) ([]Post, error) {
	cursor, err := p.coll.Find(context.Background(), Post{AuthorId: userId})
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	err = cursor.All(context.Background(), &posts)
	return posts, err
}

func (p *PostRepository) FindPostFromUserID(postId, userId string) (Post, error) {
	post := Post{}
	err := p.coll.FindOne(context.Background(), bson.M{"_id": postId, "author_id": userId}).Decode(&post)
	return post, err
}

func (p *PostRepository) FindPostFromTitle(title string) (Post, error) {
	post := Post{}
	err := p.coll.FindOne(context.Background(), bson.M{"title": title}).Decode(&post)
	return post, err
}

func (p *PostRepository) ThumbUpPost(ctx context.Context, postId string) error {
	post := Post{}
	err := p.coll.FindOne(context.Background(), bson.M{"_id": postId}).Decode(&post)
	if err != nil {
		return err
	}
	objId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	points := post.Points
	points++
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", bson.D{{"points", points}}}}
	_, err = p.coll.UpdateOne(ctx, filter, update)
	return err
}

func (p *PostRepository) ThumbDownPost(ctx context.Context, postId string) error {
	post := Post{}
	err := p.coll.FindOne(context.Background(), bson.M{"_id": postId}).Decode(&post)
	if err != nil {
		return err
	}
	points := post.Points
	points--
	filter := bson.D{{"_id", postId}}
	update := bson.D{{"$set", bson.D{{"points", points}}}}
	_, err = p.coll.UpdateOne(ctx, filter, update)
	return err
}

type Comment struct {
	PostID    string `json:"post_id" bson:"PostID,omitempty"`
	AuthorId  string `json:"author_id" bson:"Author,omitempty"`
	Body      string `json:"content" bson:"Body,omitempty"`
	CommentID string `json:"comment_id" bson:"_id,omitempty"`
	Points    string `json:"points" bson:"points,omitempty"`
}

type CommentRepository struct {
	coll *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) CommentRepository {
	return CommentRepository{coll: db.Collection("Comments")}
}

func (c *CommentRepository) AddComment(comment Comment) error {
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

func (c *CommentRepository) GetAllCommentsFromUserId(userId string) ([]Comment, error) {
	cursor, err := c.coll.Find(context.Background(), Comment{AuthorId: userId})
	if err != nil {
		return nil, err
	}
	var comments []Comment
	err = cursor.All(context.Background(), &comments)
	return comments, err
}

func (c *CommentRepository) GetAllCommentsFromPostId(postId string) ([]Comment, error) {
	cursor, err := c.coll.Find(context.Background(), Comment{PostID: postId})
	if err != nil {
		return nil, err
	}
	var comments []Comment
	err = cursor.All(context.Background(), &comments)
	return comments, err
}

func (c *CommentRepository) FindComment(postId, commentId string) (Comment, error) {
	comment := Comment{}
	err := c.coll.FindOne(context.Background(), bson.M{"PostID": postId, "CommentID": commentId}).Decode(&comment)
	return comment, err
}

func (c *CommentRepository) ThumbUpComment(ctx context.Context, postId, commentId string) error {
	comment := Comment{}
	err := c.coll.FindOne(context.Background(), bson.M{"PostID": postId, "CommentID": commentId}).Decode(&comment)
	if err != nil {
		return err
	}
	points, _ := strconv.Atoi(comment.Points)
	points++
	filter := bson.D{{"PostID", postId}, {"CommentID", commentId}}
	update := bson.D{{"$set", bson.D{{"points", strconv.Itoa(points)}}}}
	_, err = c.coll.UpdateOne(ctx, filter, update)
	return err
}

func (c *CommentRepository) ThumbDownComment(ctx context.Context, postId, commentId string) error {
	comment := Comment{}
	err := c.coll.FindOne(context.Background(), bson.M{"PostID": postId, "CommentID": commentId}).Decode(&comment)
	if err != nil {
		return err
	}
	points, _ := strconv.Atoi(comment.Points)
	points--
	filter := bson.D{{"PostID", postId}, {"CommentID", commentId}}
	update := bson.D{{"$set", bson.D{{"points", strconv.Itoa(points)}}}}
	_, err = c.coll.UpdateOne(ctx, filter, update)
	return err
}
