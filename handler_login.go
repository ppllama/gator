package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username given")
	}

	err := s.conf.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error setting user %w", err)
	}
	fmt.Printf("The user %s has been set\n", cmd.args[0])
	return nil
}