package db

import (
	"context"
	"log"
	"math/rand"
	"strconv"

	"github.com/baoduong254/gopher-social/internal/store"
)

var usernames = []string{
	"alice",
	"bob",
	"charlie",
	"dave",
	"eve",
	"frank",
	"grace",
	"heidi",
	"ivan",
	"judy",
}

var tags = []string{
	"golang",
	"programming",
	"webdev",
	"database",
	"api",
	"backend",
	"frontend",
	"devops",
	"cloud",
	"microservices",
}

var commentTexts = []string{
	"Great post!",
	"Thanks for sharing.",
	"Very informative.",
	"I learned a lot from this.",
	"Can you provide more details?",
	"I have a question about this.",
	"This is exactly what I needed.",
	"Looking forward to your next post!",
	"Keep up the good work!",
	"This is a game-changer.",
}

func randomUsername() string {
	return usernames[rand.Intn(len(usernames))] + strconv.Itoa(rand.Intn(1000))
}

func Seed(store store.Storage) error {
	ctx := context.Background()
	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			return err
		}
	}
	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			return err
		}
	}
	comments := generateComments(500, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			return err
		}
	}
	log.Printf("Seeded %d users, %d posts, and %d comments", len(users), len(posts), len(comments))
	log.Printf("Seed successful")
	return nil
}

func generateUsers(n int) []*store.User {
	users := make([]*store.User, n)
	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username: randomUsername(),
			Email:    randomUsername() + "@example.com",
			Password: "password",
		}
	}
	return users
}

func generatePosts(n int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, n)
	for i := 0; i < n; i++ {
		posts[i] = &store.Post{
			Title:   "Post " + strconv.Itoa(i+1),
			Content: "This is the content of post " + strconv.Itoa(i+1),
			UserID:  users[rand.Intn(len(users))].ID,
			Tags:    []string{tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))]},
		}
	}
	return posts
}

func generateComments(n int, posts []*store.Post, users []*store.User) []*store.Comment {
	comments := make([]*store.Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: commentTexts[rand.Intn(len(commentTexts))],
		}
	}
	return comments
}
