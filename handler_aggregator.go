package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
	"github.com/VictorHRRios/blog_aggregator/internal/feed"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Needs a time between requests")
	}
	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(time_between_reqs)
	fmt.Println("Collecing feeds every:", time_between_reqs)
	for ; ; <-ticker.C {
		fmt.Println("making a request")
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.queries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.queries.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}
	rssFeed, err := feed.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	fmt.Printf("Title: %v\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		_, err = s.queries.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: item.PubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return err
		}
		fmt.Printf("Item title: %v\n", item.Title)
	}

	return nil
}
