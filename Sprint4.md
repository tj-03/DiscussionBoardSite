# FrontEnd
- Homepage rehaul
- Login/Logout using Google auth
- Add titles to posts
- Post content and comments now show up in individual dialog box
- Dark mode

# BackEnd
- Added endpoints for searching posts/users
- Added endpoints for replying to posts
- Added endpoints for liking/upvoting posts
- Configured authorized endpoints (Firebase), allowing only those with a Google account to create posts and comments
- Updated API documentation

## Issues Worked On
- Add titles to posts.
- Search posts by title.
- Search users by username.
- Update points on posts and comments.

## New Functionality Implemented
- Posts can now contain and be searched by the title field.
- Posts now contain a points field for upvotes and downvotes.
- Comments now conntain a points field for upvotes and downvotes.
- Users can now be searched for by username.
- Users can now login/logout using Firebase


## Tests Created
- TestCommentPoints
- TestPostPoints
- TestGetUserByUsername

### Passed? Y/N
All tests passed.
