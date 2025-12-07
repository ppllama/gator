package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ppllama/gator/internal/config"
)

type state struct{
	conf *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", conf)

	session := state{
		conf: &conf,
	}

	com := commands{
		command: make(map[string]func(*state, command) error),
	}

	com.register("login", handlerLogin)

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
		log.Fatalf("error running command: %s\n", err)
	}

}