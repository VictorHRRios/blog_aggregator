package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	user, err := s.queries.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Colud not get user:\n%v", err)
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("follow command needs one argument")
	}

	feedUrl := cmd.arguments[0]
	feed, err := s.queries.GetFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Could not get feed:\n%v", err)
	}

	feedFollow, err := s.queries.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		})
	if err != nil {
		return fmt.Errorf("Could not set feed:\n%v", err)
	}

	fmt.Printf(
		"User: %v\nFeed Url: %v\n",
		feedFollow.UserName,
		feedFollow.FeedName)

	return nil
}
