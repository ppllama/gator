package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ppllama/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username given")
	}

	user, err := s.db.GetUser(context.Background(),cmd.args[0])
	if err == nil {
		return fmt.Errorf("%v already exists", cmd.args[0])
	}


	userParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	}
	
	user, err = s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error registering user: %s: %v", cmd.args[0], err)
	}
	fmt.Printf("The user %s has been registered\n", user.Name)


	err = s.conf.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting new user %w", err)
	}
	fmt.Printf("The user %s has been set\n", user.Name)


	fmt.Printf("Debug info\n%v\n", user)

	return nil
}