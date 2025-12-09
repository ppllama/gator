package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}

	if len(users) == 0 {
		fmt.Println("there are no users :(")
		return nil
	}
	
	for _, user := range(users) {
		if user == "" {
			continue
		}

		if user == s.conf.Current_user_name {
			fmt.Printf("* %s (current)\n", user)
			continue
		}

		fmt.Printf("* %s\n", user)
	}

	return nil
}