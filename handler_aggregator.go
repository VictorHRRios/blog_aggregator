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
		return fmt.Errorf("Needs an URL")
	}
	url := cmd.arguments[0]
	ctx := context.Background()
	feed, err := feed.FetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("Something went wrong when fetching url\n%v", err)
	}

	fmt.Println(feed.Channel.Title)
	fmt.Println(feed.Channel.Link)
	fmt.Println(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println(item.Description)
		fmt.Println(item.PubDate)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.queries.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Needs a name of the feed and an URL")
	}
	name := cmd.arguments[0]
	url := cmd.arguments[1]
	newFeed, err := s.queries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("newFeed: %v\n", newFeed)
	return nil
}
