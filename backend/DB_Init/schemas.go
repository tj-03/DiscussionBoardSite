package main

type User struct {
	UserId   string `json:"id" bson:"_id,omitempty"` // bson is for mongo
	Username string `json:"name" bson:"Username,omitempty"`
}

type Post struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	AuthorId string `json:"author_id" bson:"author_id,omitempty"`
	Content  string `json:"content" bson:"content,omitempty"`
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

func createTestPosts() []Post {

	return []Post{
		{
			ID:       "1",
			AuthorId: "1",
			Content:  "post1",
		},
		{
			ID:       "2",
			AuthorId: "1",
			Content:  "post2",
		},
		{
			ID:       "3",
			AuthorId: "2",
			Content:  "post3",
		},
		{
			ID:       "4",
			AuthorId: "2",
			Content:  "post4",
		},
		{
			ID:       "5",
			AuthorId: "3",
			Content:  "post5",
		},
	}

}
