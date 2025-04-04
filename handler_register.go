package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Username is required for login")
	}
	userName := cmd.arguments[0]
	ctx := context.Background()
	_, err := s.queries.CreateUser(ctx, database.CreateUserParams{
		Name:      userName,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("Could not create user:\n%v", err)
	}

	if err := s.cfg.SetUser(cmd.arguments[0]); err != nil {
		return fmt.Errorf("Could not set user:\n %v", err)
	}
	fmt.Println("User", cmd.arguments[0], "has been seet!")

	return nil
}
