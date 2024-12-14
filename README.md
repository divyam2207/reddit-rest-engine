
# Reddit Engine Project with Actor-Based Concurrency and RESTful APIs

- **Project Members**: 
- Divyam Dubey
-  Kratik Patel 

### Youtube Link:  https://youtu.be/1y66GZYYv2o

## Project Overview
This project is a simple Reddit clone built using **Go**, **Proto.Actor**, and **Gin**. It demonstrates actor-based concurrency using RESTful API to efficiently handle user interactions with posts, comments, and subreddits via REST endpoints. The system includes features like user registration, post creation, voting, and subreddit management and finally state management.

## Technologies Used
- **Go**: Programming language for backend logic
- **Proto.Actor**: Actor model library for handling concurrency
- **Gin**: Framework for building RESTful APIs
- **ProtoBuffer**: For efficient data transfer.

## Architecture Overview
The system is structured around the actor model, where the following actors are defined:
- **UserActor**: Manages user actions, including registration, posting, commenting, and voting.
- **SubredditActor**: Manages subreddit creation, membership, and posts.
- **Actor Model**: Allows for concurrent handling of actions such as user registration, joining subreddits, posting content, and voting.

## Actor Logic
### **UserActor**
- Handles:
  - User registration
  - Creating posts
  - Commenting on posts
  - Voting on posts
- Messages handled:
  - `RegisterUserMessage`: Registers a new user.
  - `PostMessage`: Handles user posts.
  - `VoteMessage`: Handles voting on posts.
  
### **SubredditActor**
- Handles:
  - Creating subreddits
  - Joining subreddits
- Messages handled:
  - `CreateSubredditMessage`: Creates a new subreddit.
  - `JoinSubredditMessage`: Adds users to subreddits.

## RESTful API Endpoints
The system exposes the following API endpoints:

1. **POST /users/register**
   - **Function**: Registers a new user.
   - **Description**: This endpoint allows a new user to sign up by providing necessary details such as username, email, and password.

2. **POST /subreddits/create**
   - **Function**: Creates a new subreddit.
   - **Description**: This endpoint enables the creation of a new subreddit, where users can submit posts, join, and interact with the community.

3. **POST /subreddits/:name/join**
   - **Function**: Allows a user to join a subreddit.
   - **Description**: By specifying the subreddit name, a user can join the subreddit, which grants them access to its posts, comments, and other activities.

4. **POST /subreddits/:name/posts**
   - **Function**: Creates a new post within a subreddit.
   - **Description**: This endpoint allows a user to create a new post under a specific subreddit, where the post will be available for other users to view and engage with.

5. **POST /subreddits/:name/posts/:postId/comments**
   - **Function**: Adds a comment to a post.
   - **Description**: This endpoint allows a user to comment on a specific post within a subreddit. It requires the post ID and the subreddit name to identify the post.

6. **POST /subreddits/:name/posts/:postId/upvote**
   - **Function**: Upvotes a post.
   - **Description**: This endpoint lets a user upvote a post, indicating approval or support for the content. It uses the `voteOnPost` function, with `true` as a parameter for upvoting.

7. **POST /subreddits/:name/posts/:postId/downvote**
   - **Function**: Downvotes a post.
   - **Description**: This endpoint allows a user to downvote a post, indicating disapproval or disagreement with the content. It uses the `voteOnPost` function, with `false` as a parameter for downvoting.

8. **GET /subreddits/:name/state**
   - **Function**: Fetches the current state of a subreddit.
   - **Description**: This endpoint retrieves the state of a specified subreddit, which could include details like the number of posts, members, or other key statistics that define the current status of the subreddit.


## Actor Interaction Flow
1. **User Registration**: A new user sends a `RegisterUserMessage` to the `UserActor`, which handles registration and returns a success message.
2. **Subreddit Creation**: A new subreddit is created by sending a `CreateSubredditMessage` to the `SubredditActor`.
3. **Joining Subreddit**: A user can join a subreddit by sending a `JoinSubredditMessage` to the relevant `SubredditActor`.
4. **Posting and Voting**: Users can post and vote on content by interacting with the appropriate actors.
## let us see this in action

## Concurrency and Scalability
- The system leverages the **actor model** to ensure that each user and subreddit interaction is handled independently and concurrently.
- The **Proto.Actor** library ensures that all user actions, such as posting or voting, are processed efficiently by separate actors without blocking other actions.

## Features Demo
1. **Register a New User**: Users can register via the `/register` endpoint.
2. **Create Subreddit**: Users can create subreddits via the `/create_subreddit` endpoint.
3. **Post Content**: Users can post content via the `/post` endpoint.
4. **Vote on Posts**: Users can vote on posts via the `/vote` endpoint.
  
All interactions are handled **concurrently** by the actor model, providing an efficient and scalable system.

## Conclusion
This project demonstrates the use of the actor-based model to efficiently handle multiple concurrent actions in a Reddit-like platform. By leveraging **Proto.Actor** and **Gin**, the system can handle user registrations, post creation, subreddit management, and voting efficiently.

Thank you for reviewing the project!
