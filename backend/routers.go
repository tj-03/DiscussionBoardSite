package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func findUser(arr []User, id string) *User {
	for _, a := range arr {
		if a.UserId == id {
			return &a
		}
	}
	return nil
}

func postExists(arr []Post, id string) bool {
	for _, a := range arr {
		if a.ID == id {
			return true
		}
	}
	return false
}

func NewMockRouter() *gin.Engine {
	posts := []Post{
		Post{
			ID:       "1",
			AuthorId: "1",
			Content:  "post1",
		},
		Post{
			ID:       "2",
			AuthorId: "1",
			Content:  "post2",
		},
		Post{
			ID:       "3",
			AuthorId: "2",
			Content:  "post3",
		},
		Post{
			ID:       "4",
			AuthorId: "2",
			Content:  "post4",
		},
	}

	comments := []Comment{
		Comment{
			PostID:    "1",
			Author:    "1",
			Body:      "comment1",
			CommentID: "1",
		},
		Comment{
			PostID:    "1",
			Author:    "1",
			Body:      "comment2",
			CommentID: "2",
		},
		Comment{
			PostID:    "2",
			Author:    "2",
			Body:      "comment3",
			CommentID: "1",
		},
	}

	users := []User{
		User{
			UserId:   "1",
			Username: "user1",
		},
		User{
			UserId:   "2",
			Username: "user2",
		},
		User{
			UserId:   "3",
			Username: "user3",
		},
	}

	router := gin.Default()
	router.Use(CorsMiddleWare())
	router.Use(static.Serve("/", static.LocalFile("../frontend/dist/discussion-board", false)))

	apiGroup := router.Group("api")

	//TODO:Abstract this into seperate function or interface/struct
	//Get posts from user
	apiGroup.GET("/user/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		if findUser(users, id) == nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		userPosts := []Post{}
		for _, post := range posts {
			if post.AuthorId == id {
				userPosts = append(userPosts, post)
			}
		}
		c.JSON(200, userPosts)
	})

	//Get all posts
	apiGroup.GET("/posts", func(c *gin.Context) {
		c.JSON(200, posts)
	})

	//Get user profile
	apiGroup.GET("/user/:id", func(c *gin.Context) {
		user := findUser(users, c.Param("id"))
		if user == nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(200, user)
	})
	return router
}
