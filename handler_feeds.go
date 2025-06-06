package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.queries.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Registered feeds:")
	for _, feed := range feeds {
		fmt.Printf("    * %v, %v, by: %v\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
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

	_, err = s.queries.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    newFeed.ID,
		})
	if err != nil {
		return err
	}
	fmt.Printf("Created by: %v\n Name: %v\n URL: %v\n", s.cfg.CurrentUserName,
		newFeed.Name,
		newFeed.Url)
	return nil
}
