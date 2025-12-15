package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ppllama/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {

	current_user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name)
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
	}
	
	if len(cmd.args) < 2 {
		return fmt.Errorf(`usage: gator addfeed <name> <url>`)
	}

	feedname := cmd.args[0]
	feedURL := cmd.args[1]

	feed := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedname,
		Url: feedURL,
		UserID: current_user.ID,
	}

	output, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("error creating new feed: %v", err)
	}

	fmt.Printf("%+v\n", output)

	if err := CreateFeedFollow(s, output.ID); err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	
	return nil
}