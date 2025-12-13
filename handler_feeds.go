package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Println("there are no feeds :(")
		return nil
	}
	
	for _, feed := range(feeds) {

		if feed.Name_2 == s.conf.Current_user_name {
			fmt.Printf("* {%s %s %s (current)}\n", feed.Name, feed.Url, feed.Name_2)
			continue
		}

		fmt.Printf("* %s\n", feed)
	}

	return nil
}