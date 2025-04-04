package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/VictorHRRios/blog_aggregator/internal/config"
	"github.com/VictorHRRios/blog_aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	cfg     *config.Config
	queries *database.Queries
}

func getCommands() commands {
	cmds := commands{
		commandName: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	return cmds
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)

	newState := state{
		cfg:     &cfg,
		queries: dbQueries,
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
