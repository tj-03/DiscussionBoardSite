package main

import (
	"context"
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
	Engine *gin.Engine
	Db     *mongo.Database
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
		panic(err)
	}
	if db == nil {
		panic("db is nil")
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

	postRepo := NewPostRepository(db)
	userRepo := NewUserRepository(db)

	s.Db = db

	engine.Use(static.Serve("/", static.LocalFile("../frontend/dist/discussion-board", false)))

	apiGroup := engine.Group("api")
	apiGroup.Use(authProvider.Middleware())

	//TODO:Abstract this into seperate function or interface/struct
	//Get posts from user
	apiGroup.GET("/user/posts/:id", func(c *gin.Context) {
		posts, err := postRepo.GetAllPostsFromUserId(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, posts)
	})

	//Get all posts
	apiGroup.GET("/posts", func(c *gin.Context) {
		posts, err := postRepo.GetAllPosts()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, posts)
	})

	//Get user profile
	apiGroup.GET("/user/:id", func(c *gin.Context) {
		user, err := userRepo.FindUser(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

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

// returns a server with a default router and mock data
func NewMockServer() Server {
	var s Server
	s.InitWithRouter(NewMockRouter())
	return s
}
