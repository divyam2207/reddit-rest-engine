package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// API Base URL
const baseURL = "http://localhost:8080"

// Utility function to make POST requests
func post(endpoint string, data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}

	resp, err := http.Post(baseURL+endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("POST request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", body)
	}

	return string(body), nil
}

// Utility function to make GET requests
func get(endpoint string) (string, error) {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return "", fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", body)
	}

	return string(body), nil
}

// Register user
func registerUser(username string) {
	data := map[string]string{"username": username}
	response, err := post("/users/register", data)
	if err != nil {
		fmt.Printf("[Register User] Failed for %s: %v\n", username, err)
		return
	}
	fmt.Printf("[Register User] Response for %s: %s\n", username, response)
}

// Create subreddit
func createSubreddit(name string) {
	data := map[string]string{"name": name}
	response, err := post("/subreddits/create", data)
	if err != nil {
		fmt.Printf("[Create Subreddit] Failed for %s: %v\n", name, err)
		return
	}
	fmt.Printf("[Create Subreddit] Response for %s: %s\n", name, response)
}

// Join subreddit
func joinSubreddit(name, username string) {
	data := map[string]string{"username": username}
	response, err := post(fmt.Sprintf("/subreddits/%s/join", name), data)
	if err != nil {
		fmt.Printf("[Join Subreddit] Failed for %s by %s: %v\n", name, username, err)
		return
	}
	fmt.Printf("[Join Subreddit] Response for %s by %s: %s\n", name, username, response)
}

// Post to subreddit
func postToSubreddit(name, author, content string) {
	data := map[string]string{"subreddit": name, "author": author, "content": content}
	response, err := post(fmt.Sprintf("/subreddits/%s/posts", name), data)
	if err != nil {
		fmt.Printf("[Post to Subreddit] Failed for %s by %s: %v\n", name, author, err)
		return
	}
	fmt.Printf("[Post to Subreddit] Response for %s by %s: %s\n", name, author, response)
}

// Comment on post
func commentOnPost(name string, postID int, author, content string) {
	data := map[string]interface{}{"subreddit": name, "post_id": postID, "author": author, "content": content}
	response, err := post(fmt.Sprintf("/subreddits/%s/posts/%d/comments", name, postID), data)
	if err != nil {
		fmt.Printf("[Comment on Post] Failed for post %d by %s: %v\n", postID, author, err)
		return
	}
	fmt.Printf("[Comment on Post] Response for post %d by %s: %s\n", postID, author, response)
}

// Upvote or Downvote post
func voteOnPost(name string, postID int, username string, upvote bool) {
	endpoint := fmt.Sprintf("/subreddits/%s/posts/%d/", name, postID)
	if upvote {
		endpoint += "upvote"
	} else {
		endpoint += "downvote"
	}
	data := map[string]interface{}{"subreddit": name, "post_id": postID, "username": username}
	response, err := post(endpoint, data)
	if err != nil {
		fmt.Printf("[Vote on Post] Failed for post %d by %s: %v\n", postID, username, err)
		return
	}
	voteType := "Upvote"
	if !upvote {
		voteType = "Downvote"
	}
	fmt.Printf("[%s on Post] Response for post %d by %s: %s\n", voteType, postID, username, response)
}

// Fetch subreddit state
func fetchSubredditState(name string) {
	response, err := get(fmt.Sprintf("/subreddits/%s/state", name))
	if err != nil {
		fmt.Printf("[Fetch Subreddit State] Failed for %s: %v\n", name, err)
		return
	}
	fmt.Printf("[Fetch Subreddit State] Response for %s: %s\n", name, response)
}

// Simulate a single client workflow
func simulateClient() {
	// Generate random data for this client
	rand.Seed(time.Now().UnixNano())
	username := fmt.Sprintf("user%d", rand.Intn(10000))
	subreddit := fmt.Sprintf("subreddit%d", rand.Intn(3))
	postContent := fmt.Sprintf("This is a random post by %s", username)
	commentContent := fmt.Sprintf("This is a comment by %s", username)

	// Perform client actions
	fmt.Println("=== Simulating Client Workflow ===")
	registerUser(username)
	createSubreddit(subreddit)
	joinSubreddit(subreddit, username)
	postToSubreddit(subreddit, username, postContent)
	commentOnPost(subreddit, 1, username, commentContent)
	voteOnPost(subreddit, 1, username, true)
	fetchSubredditState(subreddit)
	fmt.Println("=== Client Workflow Complete ===")
}

func main() {
	// Simulate a single client
	simulateClient()
}
