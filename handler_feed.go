package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nlanzo/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <feed_name> <feed_url>")
	}

	// insert the feed into the database
	newFeedRecord, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:   uuid.New(),
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Added feed %s for user %s\n", cmd.Args[1], user.Name)
	fmt.Printf("Feed ID: %s\n", newFeedRecord.ID)
	fmt.Printf("Feed Name: %s\n", newFeedRecord.Name)
	fmt.Printf("Feed URL: %s\n", newFeedRecord.Url)
	fmt.Printf("Feed User ID: %s\n", newFeedRecord.UserID)
	fmt.Printf("Feed Created At: %s\n", newFeedRecord.CreatedAt)
	fmt.Printf("Feed Updated At: %s\n", newFeedRecord.UpdatedAt)


	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    newFeedRecord.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Printf("Followed feed %s for user %s\n", cmd.Args[1], user.Name)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		fmt.Printf("Name: %s, URL: %s, User: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
