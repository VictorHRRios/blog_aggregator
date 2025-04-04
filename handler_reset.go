package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	if len(cmd.arguments) != 0 {
		return fmt.Errorf("No parameters needed for reset command")
	}

	if err := s.queries.DeleteUser(ctx); err != nil {
		return fmt.Errorf("Could not delete entries from users table:\n%v", err)
	}
	fmt.Println("Deleted entries from users table")
	return nil
}
