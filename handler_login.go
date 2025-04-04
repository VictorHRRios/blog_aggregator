package main

import (
	"context"
	"fmt"
)

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
