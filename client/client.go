package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Helper function to send POST requests
func sendPostRequest(url string, payload interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func main() {
	baseURL := "http://localhost:8080"

	// Register users
	fmt.Println("Registering users...")
	_, err := sendPostRequest(fmt.Sprintf("%s/users", baseURL), map[string]string{
		"username": "user1",
	})
	if err != nil {
		log.Fatalf("Error registering user1: %v", err)
	}
	fmt.Println("User 'user1' registered successfully.")

	_, err = sendPostRequest(fmt.Sprintf("%s/users", baseURL), map[string]string{
		"username": "user2",
	})
	if err != nil {
		log.Fatalf("Error registering user2: %v", err)
	}
	fmt.Println("User 'user2' registered successfully.")

	// Create a subreddit
	fmt.Println("Creating a subreddit...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits", baseURL), map[string]string{
		"name": "golang",
	})
	if err != nil {
		log.Fatalf("Error creating subreddit 'golang': %v", err)
	}
	fmt.Println("Subreddit 'golang' created successfully.")

	// Join a subreddit
	fmt.Println("User 'user1' joining subreddit 'golang'...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits/golang/join", baseURL), map[string]string{
		"username": "user1",
	})
	if err != nil {
		log.Fatalf("Error joining subreddit 'golang': %v", err)
	}
	fmt.Println("User 'user1' joined subreddit 'golang'.")

	// Post in the subreddit
	fmt.Println("User 'user1' posting in subreddit 'golang'...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits/golang/posts", baseURL), map[string]string{
		"author":  "user1",
		"content": "Hello, Go community!",
	})
	if err != nil {
		log.Fatalf("Error posting in subreddit 'golang': %v", err)
	}
	fmt.Println("User 'user1' posted in subreddit 'golang'.")

	// Comment on a post
	fmt.Println("User 'user2' commenting on a post in subreddit 'golang'...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits/golang/posts/1/comments", baseURL), map[string]interface{}{
		"author":          "user2",
		"content":         "Great post!",
		"parent_comment_id": 0,
	})
	if err != nil {
		log.Fatalf("Error commenting on post 1: %v", err)
	}
	fmt.Println("User 'user2' commented on post 1.")

	// Upvote the post
	fmt.Println("User 'user2' upvoting post 1 in subreddit 'golang'...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits/golang/posts/1/upvote", baseURL), map[string]string{
		"username": "user2",
	})
	if err != nil {
		log.Fatalf("Error upvoting post 1: %v", err)
	}
	fmt.Println("User 'user2' upvoted post 1.")

	// Downvote the post
	fmt.Println("User 'user2' downvoting post 1 in subreddit 'golang'...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits/golang/posts/1/downvote", baseURL), map[string]string{
		"username": "user2",
	})
	if err != nil {
		log.Fatalf("Error downvoting post 1: %v", err)
	}
	fmt.Println("User 'user2' downvoted post 1.")

	// Leave the subreddit
	fmt.Println("User 'user1' leaving subreddit 'golang'...")
	_, err = sendPostRequest(fmt.Sprintf("%s/subreddits/golang/leave", baseURL), map[string]string{
		"username": "user1",
	})
	if err != nil {
		log.Fatalf("Error leaving subreddit 'golang': %v", err)
	}
	fmt.Println("User 'user1' left subreddit 'golang'.")
}
