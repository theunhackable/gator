package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/theunhackable/gator/internal/config"
	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/handlers"
	"github.com/theunhackable/gator/internal/models"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		os.Exit(1)
	}

	database, err := sql.Open("postgres", cfg.DBUrl)
	state := models.State{
		State: &cfg,
		Db:    db.New(database),
	}
	cmds := models.Commands{
		Registered: make(map[string]func(s *models.State, c models.Command) error),
	}

	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerUsers)
	cmds.Register("agg", handlers.HandlerAgg)
	cmds.Register("addfeed", handlers.HandlerAddFeed)
	cmds.Register("feeds", handlers.HandlerFeeds)

	args := os.Args
	cmd, ok := cmds.Registered[args[1]]

	if !ok {
		fmt.Printf("Command '%s' not found", args[1])
		os.Exit(1)
	}

	if err := cmd(&state, models.Command{Name: args[1], Arguments: args}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
