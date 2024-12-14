package main

import (
	"fmt"
	"net/http"
	"time"

	pb "reddit-clone/proto" // Import the Protobuf package

	"github.com/asynkron/protoactor-go/actor"
	"github.com/gin-gonic/gin"
)

// UserActor handles user-specific actions
type UserActor struct {
	username string
}

// SubredditActor handles subreddit-specific actions
type SubredditActor struct {
	name    string
	members map[string]bool
	posts   []*pb.Post
}

// Messages for UserActor and SubredditActor
type RegisterUserMessage struct {
	Username string
}

type JoinMessage struct {
	Username string
}

type PostMessage struct {
	Author  string
	Content string
}

type CommentMessage struct {
	PostID  int32
	Author  string
	Content string
}

type VoteMessage struct {
	PostID   int32
	Username string
	Upvote   bool
}

type FetchSubredditStateMessage struct{}

// Receive processes messages for UserActor
func (u *UserActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		// Handle actor started lifecycle message
		fmt.Printf("UserActor for user '%s' started\n", u.username)
	case *RegisterUserMessage:
		fmt.Printf("User registered: %s\n", msg.Username)
		ctx.Respond(fmt.Sprintf("User %s registered successfully", msg.Username))
	default:
		fmt.Printf("Unhandled message in UserActor: %+v of type %T\n", msg, msg)
	}
}

// Receive processes messages for SubredditActor
func (s *SubredditActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		// Handle actor started lifecycle message
		fmt.Printf("SubredditActor for subreddit '%s' started\n", s.name)
	case *pb.CreateSubredditRequest:
		fmt.Printf("Subreddit created: %s\n", msg.Name)
		ctx.Respond(&pb.CreateSubredditResponse{
			Message: fmt.Sprintf("Subreddit %s created successfully", msg.Name),
		})
	case *JoinMessage:
		if s.members[msg.Username] {
			ctx.Respond(fmt.Sprintf("User %s is already a member of subreddit %s", msg.Username, s.name))
			return
		}
		s.members[msg.Username] = true
		ctx.Respond(fmt.Sprintf("User %s joined subreddit %s successfully", msg.Username, s.name))
	case *pb.PostRequest:
		post := &pb.Post{
			Id:      int32(len(s.posts) + 1),
			Author:  msg.Author,
			Content: msg.Content,
		}
		s.posts = append(s.posts, post)
		ctx.Respond(&pb.PostResponse{
			Message: fmt.Sprintf("Post created successfully by %s in subreddit %s", msg.Author, s.name),
		})
	case *CommentMessage:
		for _, post := range s.posts {
			if post.Id == msg.PostID {
				// Add comment to the post
				comment := &pb.Comment{
					Id:      int32(len(post.Comments) + 1),
					Author:  msg.Author,
					Content: msg.Content,
				}
				post.Comments = append(post.Comments, comment)
				ctx.Respond(&pb.CommentResponse{
					Message: fmt.Sprintf("Comment added by %s on post %d in subreddit %s", msg.Author, msg.PostID, s.name),
				})
				return
			}
		}
		ctx.Respond(&pb.CommentResponse{
			Message: "Post not found",
		})
	case *VoteMessage:
		for _, post := range s.posts {
			if post.Id == msg.PostID {
				if msg.Upvote {
					post.Upvotes++
					ctx.Respond(&pb.VoteResponse{
						Message: fmt.Sprintf("Upvote registered by %s on post %d in subreddit %s", msg.Username, msg.PostID, s.name),
					})
					fmt.Printf("Upvote registered by %s on post %d in subreddit %s\n", msg.Username, msg.PostID, s.name)
				} else {
					post.Downvotes++
					ctx.Respond(&pb.VoteResponse{
						Message: fmt.Sprintf("Downvote registered by %s on post %d in subreddit %s", msg.Username, msg.PostID, s.name),
					})
					fmt.Printf("Downvote registered by %s on post %d in subreddit %s\n", msg.Username, msg.PostID, s.name)
				}
				return
			}
		}
		// If post is not found
		ctx.Respond(&pb.VoteResponse{
			Message: fmt.Sprintf("Post %d not found in subreddit %s", msg.PostID, s.name),
		})
		fmt.Printf("Vote failed: Post %d not found in subreddit %s\n", msg.PostID, s.name)

	case *FetchSubredditStateMessage:
		// Gather members
		members := make([]string, 0, len(s.members))
		for member := range s.members {
			members = append(members, member)
		}

		// Respond with subreddit state
		ctx.Respond(&pb.FetchSubredditStateResponse{
			Name:    s.name,
			Members: members,
			Posts:   s.posts,
		})

	default:
		fmt.Printf("Unhandled message in SubredditActor: %+v of type %T\n", msg, msg)
	}
}

// Global variables
var (
	system          *actor.ActorSystem
	userActors      = make(map[string]*actor.PID)
	subredditActors = make(map[string]*actor.PID)
)

func init() {
	system = actor.NewActorSystem()
}

// REST handler to register a user
func registerUser(c *gin.Context) {
	var req pb.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if _, exists := userActors[req.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	props := actor.PropsFromProducer(func() actor.Actor {
		return &UserActor{username: req.Username}
	})
	rootContext := actor.NewRootContext(system, nil)
	pid := rootContext.Spawn(props)
	userActors[req.Username] = pid

	future := rootContext.RequestFuture(pid, &RegisterUserMessage{Username: req.Username}, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": result})
}

// REST handler to create a subreddit
func createSubreddit(c *gin.Context) {
	var req pb.CreateSubredditRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if _, exists := subredditActors[req.Name]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Subreddit already exists"})
		return
	}

	props := actor.PropsFromProducer(func() actor.Actor {
		return &SubredditActor{
			name:    req.Name,
			members: make(map[string]bool),
			posts:   []*pb.Post{},
		}
	})
	rootContext := actor.NewRootContext(system, nil)
	pid := rootContext.Spawn(props)
	subredditActors[req.Name] = pid

	future := rootContext.RequestFuture(pid, &req, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subreddit"})
		return
	}

	resp := result.(*pb.CreateSubredditResponse)
	c.JSON(http.StatusOK, gin.H{"response": resp.Message})
}

// REST handler to join a subreddit
func joinSubreddit(c *gin.Context) {
	subredditName := c.Param("name")
	var req pb.JoinSubredditRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if _, exists := userActors[req.Username]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found. Please register first."})
		return
	}

	pid, exists := subredditActors[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
		return
	}

	rootContext := actor.NewRootContext(system, nil)
	future := rootContext.RequestFuture(pid, &JoinMessage{Username: req.Username}, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join subreddit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": result})
}

// REST handler to post to a subreddit
func postToSubreddit(c *gin.Context) {
	var req pb.PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if _, exists := userActors[req.Author]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found. Please register first."})
		return
	}

	pid, exists := subredditActors[req.Subreddit]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
		return
	}

	rootContext := actor.NewRootContext(system, nil)
	future := rootContext.RequestFuture(pid, &req, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to post"})
		return
	}

	resp := result.(*pb.PostResponse)
	c.JSON(http.StatusOK, gin.H{"response": resp.Message})
}

func commentOnPost(c *gin.Context) {
	var req pb.CommentRequest

	// Bind the JSON request body to the Protobuf message
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate if the user exists
	if _, exists := userActors[req.Author]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found. Please register first."})
		return
	}

	// Validate if the subreddit exists
	pid, exists := subredditActors[req.Subreddit]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
		return
	}

	// Send a CommentMessage to the SubredditActor
	rootContext := actor.NewRootContext(system, nil)
	future := rootContext.RequestFuture(pid, &CommentMessage{
		PostID:  req.PostId,
		Author:  req.Author,
		Content: req.Content,
	}, 5*time.Second)

	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to comment"})
		return
	}

	// Parse the Protobuf response
	resp := result.(*pb.CommentResponse)
	c.JSON(http.StatusOK, gin.H{"response": resp.Message})
}

func voteOnPost(c *gin.Context, upvote bool) {
	var req pb.VoteRequest

	// Bind the JSON request body to the Protobuf message
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate if the user exists
	if _, exists := userActors[req.Username]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found. Please register first."})
		return
	}

	// Validate if the subreddit exists
	pid, exists := subredditActors[req.Subreddit]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
		return
	}

	// Send a VoteMessage to the SubredditActor
	rootContext := actor.NewRootContext(system, nil)
	future := rootContext.RequestFuture(pid, &VoteMessage{
		PostID:   req.PostId,
		Username: req.Username,
		Upvote:   upvote,
	}, 5*time.Second)

	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to vote on post"})
		return
	}

	// Parse the Protobuf response
	resp := result.(*pb.VoteResponse)
	c.JSON(http.StatusOK, gin.H{"response": resp.Message})
}

func fetchSubredditState(c *gin.Context) {
	subredditName := c.Param("name")

	// Validate if the subreddit exists
	pid, exists := subredditActors[subredditName]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
		return
	}

	// Send FetchSubredditStateMessage to the SubredditActor
	rootContext := actor.NewRootContext(system, nil)
	future := rootContext.RequestFuture(pid, &FetchSubredditStateMessage{}, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subreddit state"})
		return
	}

	// Parse the Protobuf response
	resp := result.(*pb.FetchSubredditStateResponse)
	c.JSON(http.StatusOK, gin.H{
		"name":    resp.Name,
		"members": resp.Members,
		"posts":   resp.Posts,
	})
}

func main() {
	r := gin.Default()
	r.POST("/users/register", registerUser)
	r.POST("/subreddits/create", createSubreddit)
	r.POST("/subreddits/:name/join", joinSubreddit)
	r.POST("/subreddits/:name/posts", postToSubreddit)
	r.POST("/subreddits/:name/posts/:postId/comments", commentOnPost)
	r.POST("/subreddits/:name/posts/:postId/upvote", func(c *gin.Context) {
		voteOnPost(c, true)
	})
	r.POST("/subreddits/:name/posts/:postId/downvote", func(c *gin.Context) {
		voteOnPost(c, false)
	})
	r.GET("/subreddits/:name/state", fetchSubredditState)

	// r.Run(":8080")

	// running on the local wifi address
	r.Run("0.0.0.0:8080")
}
