package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define routes
	r.POST("/users", registerUser)
	r.POST("/subreddits", createSubreddit)
	r.POST("/subreddits/:name/join", joinSubreddit)
	r.POST("/subreddits/:name/leave", leaveSubreddit)
	r.POST("/subreddits/:name/posts", postToSubreddit)
	r.POST("/subreddits/:name/posts/:postId/comments", commentOnPost)
	r.POST("/subreddits/:name/posts/:postId/upvote", upvotePost)
	r.POST("/subreddits/:name/posts/:postId/downvote", downvotePost)
	r.POST("/messages/reply", replyToDirectMessage)
	r.GET("/subreddits/:name/state", fetchSubredditState)

	// Start the server
	r.Run(":8080") // Run on port 8080
}
