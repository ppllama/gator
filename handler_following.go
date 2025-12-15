package main

import (
	"context"
	"fmt"

	"github.com/ppllama/gator/internal/database"
)

func handlerFollowing(s *state, _ command, currentUser database.User) error {

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}
	fmt.Println("you follow:")
	for _, feed := range(feeds) {
		fmt.Printf("%s\n", feed.Name)
	}

	return nil
}