package main

import (
	"context"
	"fmt"
	"time"

	"github.com/VictorHRRios/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerGetUsers(s *state, cmd command) error {
	ctx := context.Background()

	if len(cmd.arguments) != 0 {
		return fmt.Errorf("No need for arguments in users command")
	}
	users, err := s.queries.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("Could not get entries from users table:\n%v", err)
	}
	fmt.Println("Registered users:")
	for _, entry := range users {
		if entry.Name == s.cfg.CurrentUserName {
			fmt.Printf("    * %v (current)\n", entry.Name)
			continue
		}
		fmt.Printf("    * %v\n", entry.Name)
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	ctx := context.Background()

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Username is required for login")
	}

	userName := cmd.arguments[0]
	user, err := s.queries.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("User not in database")
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("Could not set user %v", err)
	}

	fmt.Println("User", cmd.arguments[0], "has been seet!")
	return nil
}

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
