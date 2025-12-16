package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ppllama/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	if len(cmd.args) >= 1 {
		tempLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("error parsing limit: %v", err)
		}
		limit = tempLimit
	}
	

	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	}
	
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return fmt.Errorf("error getting posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("there are no posts :(")
		return nil
	}
	
	for i, post := range(posts) {
		if post.Title.Valid {
			fmt.Printf("Post %d: %s\n", i + 1, post.Title.String)
			fmt.Println()
		}
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
			fmt.Println()
		}
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println()
		if post.PublishedAt.Valid {
			fmt.Printf("Date: %s\n", post.PublishedAt.Time)
		}
	}

	return nil
}