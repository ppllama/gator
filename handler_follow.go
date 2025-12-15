package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ppllama/gator/internal/database"
)

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	
	if len(cmd.args) == 0 {
		return fmt.Errorf("no links given")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	if err := CreateFeedFollow(s, feed.ID, currentUser.ID); err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	return nil
}

func CreateFeedFollow(s *state, feedId uuid.UUID, userID uuid.UUID) error {
// Used in add_feed

	feedFollowParams := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		UserID:		userID,
		FeedID:		feedId,
	}

	feeds, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Printf("FeedName: %s\nnow followed by\nUserName: %s\n", feeds.FeedName, feeds.UserName)
	return nil
}