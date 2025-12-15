package main

import (
	"context"
	"fmt"

	"github.com/ppllama/gator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		current_user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name)
			if err != nil {
				return fmt.Errorf("error getting current user: %v", err)
			}
		if err := handler(s, c, current_user); err != nil {
			return fmt.Errorf("error calling handler: %v", err)
		}
		return nil
	}
}