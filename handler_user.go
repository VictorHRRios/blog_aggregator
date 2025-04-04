package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Username is required for login")
	}
	if err := s.cfg.SetUser(cmd.arguments[0]); err != nil {
		return fmt.Errorf("Could not set user %v", err)
	}
	fmt.Println("User", cmd.arguments[0], "has been seet!")
	return nil
}
