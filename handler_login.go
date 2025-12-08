package main

import (
	"fmt"
	"context"
)

func handlerLogin(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username given")
	}

	user, err := s.db.GetUser(context.Background(),cmd.args[0])
	if err != nil {
		return fmt.Errorf("user not found. register using gator register <name>")
	}

	err = s.conf.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user %w", err)
	}
	fmt.Printf("The user %s has been set\n", cmd.args[0])
	return nil
}