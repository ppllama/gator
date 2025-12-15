package main

import (
	"context"
	"fmt"

)

func handlerFollowing(s *state, _ command) error {

	user, err := s.db.GetUser(context.Background(), s.conf.Current_user_name)
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}
	fmt.Println("you follow:")
	for _, feed := range(feeds) {
		fmt.Printf("%s\n", feed.Name)
	}

	return nil
}