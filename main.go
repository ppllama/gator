package main

import (
	"log"
	"os"
	"database/sql"
	"github.com/ppllama/gator/internal/database"
	"github.com/ppllama/gator/internal/config"
)

import _ "github.com/lib/pq"

type state struct{
	db		*database.Queries
	conf	*config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", conf.Db_url)
	if err != nil {
		log.Fatalf("error opening postgres: %v", err)
	}

	dbQueries := database.New(db)
	session := state{
		db:		dbQueries,
		conf:	&conf,
	}

	com := commands{
		command: make(map[string]func(*state, command) error),
	}

	com.register("login", handlerLogin)
	com.register("register", handlerRegister)
	com.register("reset", handlerReset)
	com.register("users", handlerUsers)
	com.register("agg", handlerAgg)
	com.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	com.register("feeds", handlerFeeds)
	com.register("follow", middlewareLoggedIn(handlerFollow))
	com.register("following", middlewareLoggedIn(handlerFollowing))
	com.register("unfollow", middlewareLoggedIn(handlerUnfollow))


	args := os.Args
	if len(args) < 2 {
		log.Fatalf("gator requires atleast one command")
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = com.run(&session, cmd)
	if err != nil {
		log.Fatalf("error running command: %s", err)
	}

}