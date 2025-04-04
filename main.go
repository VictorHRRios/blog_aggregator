package main

import (
	"log"
	"os"

	"github.com/VictorHRRios/blog_aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func getCommands() commands {
	cmds := commands{
		commandName: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	return cmds
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	newState := state{
		cfg: &cfg,
	}
	if len(os.Args) < 2 {
		log.Fatal("Usage blog_aggregator [command]")
	}

	commands := getCommands()
	cmdName := os.Args[1]
	cmdArguments := os.Args[2:]

	err = commands.run(&newState, command{cmdName, cmdArguments})
	if err != nil {
		log.Fatal(err)
	}

}
