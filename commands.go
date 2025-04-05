package main

import "errors"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandName map[string]func(*state, command) error
}

func getCommands() commands {
	cmds := commands{
		commandName: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	return cmds
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandName[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	funcCommand, ok := c.commandName[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}
	return funcCommand(s, cmd)
}
