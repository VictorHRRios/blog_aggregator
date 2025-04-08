package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	limit = 2

	if len(cmd.arguments) == 1 {
		limitByUser, err := strconv.ParseInt(cmd.arguments[0], 10, 32)
		if err != nil {
			fmt.Println("error")
			return err
		}
		limit = max(int32(limitByUser), limit)
	}

	posts, err := s.queries.GetPostForUser(context.Background(), database.GetPostForUserParams{
		ID:    user.ID,
		Limit: limit,
	})

	if err != nil {
		fmt.Println("error")
		return err
	}
	for _, post := range posts {
		println(post.Title)
		println(post.PublishedAt)
		println(post.Description)
		println(post.Url)
	}

	return nil
}
