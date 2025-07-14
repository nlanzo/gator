package main

import (
	"context"
	"fmt"

	"github.com/nlanzo/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <feed_name> <feed_url>")
	}
	userName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}

	// fetch the feed into the RSSFeed struct
	feed, err := fetchFeed(context.Background(), cmd.Args[1])
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	// insert the feed into the database
	newFeedRecord, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   feed.Channel.Title,
		Url:    cmd.Args[1],
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Added feed %s for user %s\n", cmd.Args[1], userName)
	fmt.Printf("Feed ID: %s\n", newFeedRecord.ID)
	fmt.Printf("Feed Name: %s\n", newFeedRecord.Name)
	fmt.Printf("Feed URL: %s\n", newFeedRecord.Url)
	fmt.Printf("Feed User ID: %s\n", newFeedRecord.UserID)
	fmt.Printf("Feed Created At: %s\n", newFeedRecord.CreatedAt)
	fmt.Printf("Feed Updated At: %s\n", newFeedRecord.UpdatedAt)

	return nil
}