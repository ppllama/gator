package main

import (
	"fmt"
	"context"
	"github.com/ppllama/gator/internal/rss_feed"
)

func handlerAgg(_ *state, _ command) error {
	
	// if len(cmd.args) == 0 {
	// 	return fmt.Errorf("no links given")
	// }

	// feedURL := cmd.args[0]
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rssfeed.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	fmt.Printf("%+v\n", *feed)

	return nil
}