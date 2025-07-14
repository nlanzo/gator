package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, _ command) error {
	// if len(cmd.Args) < 1 {
	// 	return fmt.Errorf("usage: agg <feed_url>")
	// }

	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	fmt.Printf("Fetched %d items from %s\n", len(feed.Channel.Item), feedURL)

	fmt.Println(feed)
	// for _, item := range feed.Channel.Item {
	// 	fmt.Printf("* %s\n", item.Title)
	// }

	return nil
}
