package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, _ command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting all users: %v", err)
	}

	println("successfully deleted all users")
	return nil
}