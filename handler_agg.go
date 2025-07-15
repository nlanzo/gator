package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nlanzo/gator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: agg <time_between_requests>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid time between requests: %v", err)
	}
	fmt.Printf("Collecting feeds every %s...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}	

}

func scrapeFeeds(s *state) {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("failed to get next feed to fetch: %v", err)
		return
	}

	fmt.Printf("Fetching %s\n", nextFeed.Name)

	scrapeFeed(s.db, nextFeed)

}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("failed to mark feed %s as fetched: %v", feed.ID, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("failed to fetch feed %s: %v", feed.Name, err)
		return
	}



	for _, item := range feedData.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}

	fmt.Printf("Fetched %d items from %s\n", len(feedData.Channel.Item), feed.Name)

}

