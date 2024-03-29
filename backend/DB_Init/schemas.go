package main

type User struct {
	UserId   string `json:"id" bson:"_id,omitempty"` // bson is for mongo
	Username string `json:"name" bson:"Username,omitempty"`
}

type Post struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title,omitempty"`
	AuthorId string `json:"author_id" bson:"author_id,omitempty"`
	Content  string `json:"content" bson:"content,omitempty"`
	Points   string `json:"points" bson:"points,omitempty"`
}

type Comment struct {
	PostID    string `json:"post_id" bson:"PostID,omitempty"`
	Author    string `json:"author" bson:"Author,omitempty"`
	Body      string `json:"content" bson:"Body,omitempty"`
	CommentID string `json:"comment_id" bson:"CommentID,omitempty"`
	Points    string `json:"points" bson:"points,omitempty"`
}

func createTestUsers() []User {
	return []User{
		{
			UserId:   "1",
			Username: "user1",
		},
		{
			UserId:   "2",
			Username: "user2",
		},
		{
			UserId:   "3",
			Username: "user3",
		},
		{
			UserId:   "4",
			Username: "user4",
		},
		{
			UserId:   "5",
			Username: "user5",
		},
	}

}

func createTestComments() []Comment {
	return []Comment{
		{
			PostID:    "1",
			Author:    "1",
			Body:      "comment1",
			CommentID: "1",
			Points:    "0",
		},
		{
			PostID:    "1",
			Author:    "1",
			Body:      "comment2",
			CommentID: "2",
			Points:    "0",
		},
		{
			PostID:    "2",
			Author:    "2",
			Body:      "comment3",
			CommentID: "1",
			Points:    "0",
		},
	}
}

func createTestPosts() []Post {

	return []Post{
		{
			ID:       "1",
			Title:    "Title1",
			AuthorId: "1",
			Content:  "post1",
			Points:   "0",
		},
		{
			ID:       "2",
			Title:    "Title2",
			AuthorId: "1",
			Content:  "post2",
			Points:   "0",
		},
		{
			ID:       "3",
			Title:    "Title3",
			AuthorId: "2",
			Content:  "post3",
			Points:   "0",
		},
		{
			ID:       "4",
			Title:    "Title4",
			AuthorId: "2",
			Content:  "post4",
			Points:   "0",
		},
		{
			ID:       "5",
			Title:    "Title5",
			AuthorId: "3",
			Content:  "post5",
			Points:   "0",
		},
	}

}
