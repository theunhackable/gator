package main

import (
	"os"

	"github.com/theunhackable/gator/internal/config"
	"github.com/theunhackable/gator/internal/handlers"
	"github.com/theunhackable/gator/internal/models"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		os.Exit(1)
	}
	state := models.State{
		State: &cfg,
	}

	cmd := models.Command{}
	cmd.Arguments = os.Args
	cmd.Name = "login"

	cmds := models.Commands{
		Registered: make(map[string]func(s *models.State, c models.Command) error),
	}
	cmds.Register("login", handlers.HandlerLogin)

	x, ok := cmds.Registered["login"]
	if !ok {
		os.Exit(1)
	}
	if err := x(&state, cmd); err != nil {
		os.Exit(1)
	}
}
