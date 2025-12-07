package main

import "fmt"

type command struct{
	name	string
	args	[]string
}

type commands struct{
	command		map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.command[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command")
	}
	err := function(s, cmd)
	if err!= nil {
		return fmt.Errorf("error running command %s: %w", cmd.name, err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.command[name] = f
}