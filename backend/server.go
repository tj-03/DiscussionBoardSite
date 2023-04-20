package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
)

type Server struct {
	Engine      *gin.Engine
	Db          *mongo.Database
	UserRepo    UserRepository
	PostRepo    PostRepository
	CommentRepo CommentRepository
}

func CorsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func getDbConn(dbUrl string) (*mongo.Database, error) {
	//Instantiates the client and connection location
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

	return client.Database(os.Getenv("DB_NAME")), nil
}

func (s *Server) Init(dbUrl string) {
	engine := gin.Default()
	db, err := getDbConn(dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal("db is nil")
	}

	//Set up firebase auth
	opt := option.WithCredentialsFile("./firebaseKeys.json")
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	firebaseAuthClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	authProvider := NewFirebaseAuthProvider(firebaseAuthClient, db)

	s.PostRepo = NewPostRepository(db)
	s.UserRepo = NewUserRepository(db)
	s.CommentRepo = NewCommentRepository(db)

	s.Db = db

	engine.Use(CorsMiddleWare())
	engine.Use(static.Serve("/", static.LocalFile("../frontend/dist/discussion-board", false)))

	authApiGroup := engine.Group("api")
	noAuthApiGroup := engine.Group("api")
	authApiGroup.Use(authProvider.Middleware())
	//Get posts from user
	noAuthApiGroup.GET("/user/posts/:id", s.GetPostsFromUserId)
	//Get all posts
	noAuthApiGroup.GET("/posts", s.GetAllPosts)
	//Get user profile
	noAuthApiGroup.GET("/user/:id", s.GetUser)
	authApiGroup.POST("/post", s.CreatePost)
	//Get comments
	noAuthApiGroup.GET("/post/comments/:id", s.GetCommentsFromPostId)
	authApiGroup.POST("/post/comment", s.CreateComment)
	authApiGroup.POST("/post/like", s.LikePost)
	s.Engine = engine
}

func (s *Server) InitWithRouter(router *gin.Engine) {
	s.Engine = router
}

func (s *Server) Run(port string) error {
	return s.Engine.Run("localhost:" + port)
}
func NewServer(dbUrl string) Server {
	var s Server
	s.Init(dbUrl)
	return s
}

func (s *Server) GetUser(c *gin.Context) {
	user, err := s.UserRepo.FindUser(c.Param("id"))
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
func (s *Server) GetPostsFromUserId(c *gin.Context) {
	posts, err := s.PostRepo.GetAllPostsFromUserId(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, posts)
}

func (s *Server) GetAllPosts(c *gin.Context) {
	posts, err := s.PostRepo.GetAllPosts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, posts)
}

func (s *Server) CreatePost(c *gin.Context) {
	var post Post
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(post.AuthorId) == 0 {
		c.JSON(400, gin.H{"error": "post failed, no author id"})
		return
	}
	if len(post.Content) == 0 {
		c.JSON(400, gin.H{"error": "post failed, no content"})
		return
	}

	if len(post.Title) == 0 {
		post.Title = "Untitled"
	}
	user, err := s.UserRepo.FindUser(post.AuthorId)
	println(user.UserId)
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(404, gin.H{"error": "post failed, user not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = s.PostRepo.AddPost(post)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (s *Server) LikePost(c *gin.Context) {
	postId := c.Param("id")
	if len(postId) == 0 {
		c.JSON(400, gin.H{"error": "post failed, no post id"})
		return
	}
	_, err := s.PostRepo.FindPost(postId)
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(404, gin.H{"error": "post failed, post not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	s.PostRepo.ThumbUpPost(context.Background(), postId)
}
func (s *Server) CreateComment(c *gin.Context) {
	var comment Comment
	err := c.BindJSON(&comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(comment.AuthorId) == 0 {
		c.JSON(400, gin.H{"error": "comment failed, no author id"})
		return
	}
	if len(comment.Body) == 0 {
		c.JSON(400, gin.H{"error": "comment failed, no content"})
		return
	}
	if len(comment.PostID) == 0 {
		c.JSON(400, gin.H{"error": "comment failed, no post id"})
		return
	}
	_, err = s.UserRepo.FindUser(comment.AuthorId)
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(404, gin.H{"error": "comment failed, user not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = s.PostRepo.FindPost(comment.PostID)
	if errors.Is(err, mongo.ErrNoDocuments) {
		c.JSON(404, gin.H{"error": "comment failed, post not found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = s.CommentRepo.AddComment(comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (s *Server) GetCommentsFromPostId(c *gin.Context) {
	postId := c.Param("id")
	if len(postId) == 0 {
		c.JSON(400, gin.H{"error": "no post id"})
		return
	}
	comments, err := s.CommentRepo.GetAllCommentsFromPostId(postId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, comments)
}

// returns a server with a default router and mock data
func NewMockServer() Server {
	var s Server
	s.InitWithRouter(NewMockRouter())
	return s
}
