package main

import (
	"fmt"
	"log"

	"github.com/ppllama/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", conf)

	err = conf.SetUser("lane")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	conf, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", conf)
}