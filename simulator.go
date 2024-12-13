package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Simulator struct {
	baseURL       string
	numUsers      int
	numSubreddits int
	users         []string
	subreddits    []string
	mu            sync.Mutex
}

// NewSimulator creates a new simulator instance
func NewSimulator(baseURL string, numUsers, numSubreddits int) *Simulator {
	return &Simulator{
		baseURL:       baseURL,
		numUsers:      numUsers,
		numSubreddits: numSubreddits,
		users:         []string{},
		subreddits:    []string{},
	}
}

// Helper function to send POST requests
func sendPostRequest(url string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}
	return nil
}

// Register all users
func (sim *Simulator) registerUsers() {
	for i := 1; i <= sim.numUsers; i++ {
		username := fmt.Sprintf("user%d", i)
		err := sendPostRequest(fmt.Sprintf("%s/users", sim.baseURL), map[string]string{
			"username": username,
		})
		if err != nil {
			log.Printf("Failed to register user '%s': %v", username, err)
		} else {
			sim.mu.Lock()
			sim.users = append(sim.users, username)
			sim.mu.Unlock()
			log.Printf("User '%s' registered successfully", username)
		}
	}
}

// Create subreddits
func (sim *Simulator) createSubreddits() {
	for i := 1; i <= sim.numSubreddits; i++ {
		subredditName := fmt.Sprintf("subreddit%d", i)
		err := sendPostRequest(fmt.Sprintf("%s/subreddits", sim.baseURL), map[string]string{
			"name": subredditName,
		})
		if err != nil {
			log.Printf("Failed to create subreddit '%s': %v", subredditName, err)
		} else {
			sim.mu.Lock()
			sim.subreddits = append(sim.subreddits, subredditName)
			sim.mu.Unlock()
			log.Printf("Subreddit '%s' created successfully", subredditName)
		}
	}
}

// Simulate user activities
func (sim *Simulator) simulateUserActivity(username string, wg *sync.WaitGroup) {
	defer wg.Done()
	rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(len(username))))

	for i := 0; i < 10; i++ { // Limit actions to 10 for simplicity
		action := rng.Intn(6) // Random action (0-5)
		subreddit := sim.subreddits[rng.Intn(len(sim.subreddits))]

		switch action {
		case 0: // Join subreddit
			err := sendPostRequest(fmt.Sprintf("%s/subreddits/%s/join", sim.baseURL, subreddit), map[string]string{
				"username": username,
			})
			if err != nil {
				log.Printf("User '%s' failed to join subreddit '%s': %v", username, subreddit, err)
			} else {
				log.Printf("User '%s' joined subreddit '%s'", username, subreddit)
			}

		case 1: // Post in subreddit
			err := sendPostRequest(fmt.Sprintf("%s/subreddits/%s/posts", sim.baseURL, subreddit), map[string]string{
				"author":  username,
				"content": fmt.Sprintf("This is a post by %s", username),
			})
			if err != nil {
				log.Printf("User '%s' failed to post in subreddit '%s': %v", username, subreddit, err)
			} else {
				log.Printf("User '%s' posted in subreddit '%s'", username, subreddit)
			}

		case 2: // Comment on a post
			postID := rng.Intn(5) + 1 // Random post ID
			err := sendPostRequest(fmt.Sprintf("%s/subreddits/%s/posts/%d/comments", sim.baseURL, subreddit, postID), map[string]interface{}{
				"author":          username,
				"content":         fmt.Sprintf("Comment by %s", username),
				"parent_comment_id": 0,
			})
			if err != nil {
				log.Printf("User '%s' failed to comment on post %d in subreddit '%s': %v", username, postID, subreddit, err)
			} else {
				log.Printf("User '%s' commented on post %d in subreddit '%s'", username, postID, subreddit)
			}

		case 3: // Upvote a post
			postID := rng.Intn(5) + 1 // Random post ID
			err := sendPostRequest(fmt.Sprintf("%s/subreddits/%s/posts/%d/upvote", sim.baseURL, subreddit, postID), map[string]string{
				"username": username,
			})
			if err != nil {
				log.Printf("User '%s' failed to upvote post %d in subreddit '%s': %v", username, postID, subreddit, err)
			} else {
				log.Printf("User '%s' upvoted post %d in subreddit '%s'", username, postID, subreddit)
			}

		case 4: // Downvote a post
			postID := rng.Intn(5) + 1 // Random post ID
			err := sendPostRequest(fmt.Sprintf("%s/subreddits/%s/posts/%d/downvote", sim.baseURL, subreddit, postID), map[string]string{
				"username": username,
			})
			if err != nil {
				log.Printf("User '%s' failed to downvote post %d in subreddit '%s': %v", username, postID, subreddit, err)
			} else {
				log.Printf("User '%s' downvoted post %d in subreddit '%s'", username, postID, subreddit)
			}

		case 5: // Leave subreddit
			err := sendPostRequest(fmt.Sprintf("%s/subreddits/%s/leave", sim.baseURL, subreddit), map[string]string{
				"username": username,
			})
			if err != nil {
				log.Printf("User '%s' failed to leave subreddit '%s': %v", username, subreddit, err)
			} else {
				log.Printf("User '%s' left subreddit '%s'", username, subreddit)
			}
		}
	}
}

// Run the simulation
func (sim *Simulator) runSimulation() {
	sim.registerUsers()
	sim.createSubreddits()

	var wg sync.WaitGroup
	for _, username := range sim.users {
		wg.Add(1)
		go sim.simulateUserActivity(username, &wg)
	}
	wg.Wait()
}

func main() {
	baseURL := "http://localhost:8080"
	simulator := NewSimulator(baseURL, 100, 10) // 100 users, 10 subreddits
	simulator.runSimulation()
}
