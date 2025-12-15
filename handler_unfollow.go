package main

import (
	"context"
	"fmt"

	"github.com/ppllama/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	
	if len(cmd.args) == 0 {
		return fmt.Errorf("no links given")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	params := database.DeleteFeedFollowForUserParams{
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	if err := s.db.DeleteFeedFollowForUser(context.Background(), params); err != nil {
		return fmt.Errorf("error deleting feed follow: %v", err)
	}

	fmt.Printf("unfollowed %s\n", feed.Name)

	return nil
}