package db

import (
	"context"
	"database/sql"
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

func Seed(store store.Storage, db *sql.DB) error {
	ctx := context.Background()
	users, err := generateUsers(100)
	if err != nil {
		return err
	}
	tx, _ := db.BeginTx(ctx, nil)
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
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
	followers := generateFollowers(300, users)
	for _, pair := range followers {
		if err := store.Followers.Follow(ctx, pair[0], pair[1]); err != nil {
			log.Printf("Error creating follower relationship: %v", err)
		}
	}
	log.Printf("Seeded %d users, %d posts, %d comments, and %d follower relationships", len(users), len(posts), len(comments), len(followers))
	log.Printf("Seed data generation complete")
	return nil
}

func generateUsers(n int) ([]*store.User, error) {
	users := make([]*store.User, n)
	runID := strconv.FormatInt(rand.Int63(), 10)
	for i := 0; i < n; i++ {
		username := randomUsername() + "_" + runID + "_" + strconv.Itoa(i)
		user := &store.User{
			Username: username,
			Email:    username + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}

		if err := user.Password.Set("password"); err != nil {
			return nil, err
		}

		users[i] = user
	}
	return users, nil
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

func generateFollowers(n int, users []*store.User) [][2]int64 {
	followers := make([][2]int64, n)
	for i := 0; i < n; i++ {
		followers[i] = [2]int64{
			users[rand.Intn(len(users))].ID,
			users[rand.Intn(len(users))].ID,
		}
	}
	return followers
}
