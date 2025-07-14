package main

import (
	"context"
	"fmt"

	"github.com/nlanzo/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: followfeed <feed_url>")
	}

	feedURL := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed: %v", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("Followed feed %s for user %s\n", feed.Name, user.Name)


	return nil
}

func handlerListFollowedFeeds(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Printf("No feeds followed by %s\n", s.cfg.CurrentUserName)
		return nil
	}

	fmt.Printf("Found %d feeds followed by %s:\n", len(feeds), s.cfg.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("Feed: %s\n", feed.FeedName)
	}

	return nil
}