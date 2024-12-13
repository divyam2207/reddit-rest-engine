package main

import (
	"net/http"
	"sync"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
	pb "reddit-clone/proto"
	
)

// Shared data structures
var (
	users      = make(map[string]*User)
	subreddits = make(map[string]*Subreddit)
	mu         sync.Mutex
)

// Internal data models
type User struct {
	Username string
	Karma    int
	Inbox []pb.DirectMessageRequest

}

type Subreddit struct {
	Name    string
	Members map[string]bool
	Posts   []*pb.Post
}

// REST Handlers

// Register a new user
func registerUser(c *gin.Context) {
	var req pb.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[req.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}
	users[req.Username] = &User{Username: req.Username, Karma: 0, Inbox: []pb.DirectMessageRequest{}}
	c.JSON(http.StatusOK, pb.RegisterUserResponse{Message: "User registered"})
}

// Create a new subreddit
func createSubreddit(c *gin.Context) {
	var req pb.CreateSubredditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := subreddits[req.Name]; exists {
		c.JSON(http.StatusConflict, gin.H{"message": "Subreddit already exists"})
		return
	}
	subreddits[req.Name] = &Subreddit{
		Name:    req.Name,
		Members: make(map[string]bool),
		Posts:   []*pb.Post{},
	}
	log.Printf("Current subreddits: %v", subreddits)
	c.JSON(http.StatusOK, pb.CreateSubredditResponse{Message: "Subreddit created"})
}

// Reply to a direct message
func replyToDirectMessage(c *gin.Context) {
	var req pb.DirectMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	recipient, exists := users[req.To]
	if !exists {
		c.JSON(http.StatusNotFound, pb.DirectMessageResponse{Message: "Recipient not found"})
		return
	}

	recipient.Inbox = append(recipient.Inbox, pb.DirectMessageRequest{
		From:    req.From,
		To:      req.To,
		Content: req.Content,
	})
	
	c.JSON(http.StatusOK, pb.DirectMessageResponse{Message: "Reply sent successfully"})
}

// Join a subreddit
func joinSubreddit(c *gin.Context) {
	var req pb.JoinSubredditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	subredditName := c.Param("name")
	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, pb.JoinSubredditResponse{Success: false, Message: "Subreddit not found"})
		return
	}

	if _, alreadyMember := subreddit.Members[req.Username]; alreadyMember {
		c.JSON(http.StatusConflict, pb.JoinSubredditResponse{Success: false, Message: "User is already a member"})
		return
	}

	subreddit.Members[req.Username] = true
	c.JSON(http.StatusOK, pb.JoinSubredditResponse{Success: true, Message: "Joined subreddit successfully"})
}

// Post to a subreddit
func postToSubreddit(c *gin.Context) {
	var req pb.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	subredditName := c.Param("name")
	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, pb.PostResponse{Message: "Subreddit not found"})
		return
	}

	post := &pb.Post{
		Id:      int32(len(subreddit.Posts) + 1),
		Author:  req.Author,
		Content: req.Content,
	}
	subreddit.Posts = append(subreddit.Posts, post)
	c.JSON(http.StatusOK, pb.PostResponse{Message: "Post created"})
}

// Comment on a post
func commentOnPost(c *gin.Context) {
	var req pb.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	
	subredditName := c.Param("name")
	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, pb.CommentResponse{Message: "Subreddit not found"})
		return
	}

	postId_param := c.Param("postId")
	postId_conv, err := strconv.ParseInt(postId_param, 10, 32)
	if err == nil {
		for _, post := range subreddit.Posts {
			if post.Id == int32(postId_conv) {
				comment := &pb.Comment{
					Id:      int32(len(post.Comments) + 1),
					Author:  req.Author,
					Content: req.Content,
				}
				if req.ParentCommentId == 0 {
					post.Comments = append(post.Comments, comment)
				} else {
					for _, parent := range post.Comments {
						if parent.Id == req.ParentCommentId {
							parent.Replies = append(parent.Replies, comment)
						}
					}
				}
				c.JSON(http.StatusOK, pb.CommentResponse{Message: "Comment added"})
				return
			}
		}
	}
	
	c.JSON(http.StatusNotFound, pb.CommentResponse{Message: "Post not found"})
}

// Upvote a post
func upvotePost(c *gin.Context) {
	var req pb.UpvotePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	subredditName := c.Param("name")
	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, pb.UpvotePostResponse{Message: "Subreddit not found"})
		return
	}

	postId_param := c.Param("postId")
	postId_conv, err := strconv.ParseInt(postId_param, 10, 32)
	if err == nil {
		for _, post := range subreddit.Posts {
			if post.Id == int32(postId_conv) {
				post.Upvotes++
				c.JSON(http.StatusOK, pb.UpvotePostResponse{Message: "Upvote successful"})
				return
			}
		}
	}
	c.JSON(http.StatusNotFound, pb.UpvotePostResponse{Message: "Post not found"})
}

// Leave a subreddit
func leaveSubreddit(c *gin.Context) {
	var req pb.LeaveSubredditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Retrieve the subreddit
	subredditName := c.Param("name")
	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, pb.LeaveSubredditResponse{Success: false, Message: "Subreddit not found"})
		return
	}

	// Check if the user is a member
	if _, isMember := subreddit.Members[req.Username]; !isMember {
		c.JSON(http.StatusConflict, pb.LeaveSubredditResponse{Success: false, Message: "User is not a member of the subreddit"})
		return
	}

	// Remove the user from the subreddit
	delete(subreddit.Members, req.Username)
	c.JSON(http.StatusOK, pb.LeaveSubredditResponse{Success: true, Message: "Left subreddit successfully"})
}

// Downvote a post
func downvotePost(c *gin.Context) {
	var req pb.DownvotePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	
	subredditName := c.Param("name")
	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, pb.DownvotePostResponse{Message: "Subreddit not found"})
		return
	}

	postId_param := c.Param("postId")
	postId_conv, err := strconv.ParseInt(postId_param, 10, 32)
	if err == nil {
		for _, post := range subreddit.Posts {
			if post.Id == int32(postId_conv) {
				post.Downvotes++
				c.JSON(http.StatusOK, pb.DownvotePostResponse{Message: "Downvote successful"})
				return
			}
		}
	}
	
	c.JSON(http.StatusNotFound, pb.DownvotePostResponse{Message: "Post not found"})
}

// Fetch subreddit state
func fetchSubredditState(c *gin.Context) {
	subredditName := c.Param("name")

	mu.Lock()
	defer mu.Unlock()

	subreddit, exists := subreddits[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
		return
	}

	var members []string
	for member := range subreddit.Members {
		members = append(members, member)
	}

	c.JSON(http.StatusOK, pb.FetchSubredditStateResponse{
		Name:    subreddit.Name,
		Members: members,
		Posts:   subreddit.Posts,
	})
}

// Main function to run the server
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
