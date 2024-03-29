# Work Completed


## Backend
- Added retrieval functions for posts and users given different parameters.
- Added functions to add or retreive comments.
- Added/modified unit tests to guarantee full backend functionality.
- Created server setup documentation
- Created Swagger API documentation
- Updated endpoints to be consistent with Swagger documentation and ease frontend use 

## Frontend
- Created login page.
- Created sign-up page.
- Created thumb-up and thumb-down buttons.
- Created share button.
- Created post format.
- Set up framework to ease integrating backend with frontend in future sprints

# Backend Unit Tests
All backend unit test run using a local mongodb database with sample users, posts, and comments.

- TestGetAllPosts: tests the GetAllPosts() function and whether it returns ALL posts.
- TestGetExistingPost: tests whether the FindPost() function returns the right post, given that post exists.
- TestPostDoesNotExist: tests whether the appropriate response is given when FindPost() tries to find a non-existant post.
- TestAddPost: tests the AddPost() function and whether the post is added to the database.
- TestGetPostFromUser: tests the FindPostFromUserID() function and whether the correct post is assigned to a specified user.
- TestGetAllPostsFromUser: test the GetAllPostsFromUserID function and whether all relevant user's posts are returned.
- TestGetAllUser: Tests the GetAllUsers() function and whether all users are returned.
- TestGetUser: Tests FindUser() and whether the correct user is returned given a user exists.
- TestGetUserThatDoesNotExist: Tests FindUser() and its response to a non-existant user.
- TestAddUser: Tests CreateNewUser() and whether the user is added to the DB.
- TestGetAllComments: Tests GetAllComments() and whether all comments are returned.
- TestGetAllCommentsFromUser: Tests GetAllCommentsFromUserId() and whether the correct number of comments are returned.
- TestGetAllCommentsFromPost: Tests GetAllCommentsFromPostID() and whether the correct number of comments are returned.
- TestGetExistingComment: Tests FindComment() and whether the correct comment is returned given that the comment exists.
- TestCommentThatDoesNotExist: Test FindComment() and its response to a non-existant comment.
- TestAddComment: Tests AddComment() and whether the comment is added to the DB.

All test pass.

# Frontend Unit Tests

- Signup: Ensures login button leads to the sign-up button which leads to the proper sign-up page.
- HomeCall: Makes sure home page properly fetches posts from API.

# Backend API Documentation
https://github.com/tj-03/DiscussionBoardSite/tree/main/backend/docs/docs.html
# Plan for Next Sprint
- Add title to posts.
- Add 'thumb up/down' to posts and comments.
- Add 'points' to users.
- Integrate sign-in page with Firebase authentication.
