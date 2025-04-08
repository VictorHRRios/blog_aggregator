package main

import (
	"context"
	"fmt"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.queries.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Could not get user:\n%v", err)
		}
		return handler(s, cmd, user)
	}
}
