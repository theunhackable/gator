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

	cmdLogin := models.Command{}
	cmdLogin.Arguments = os.Args
	cmdLogin.Name = "login"

	cmdRegister := models.Command{}
	cmdRegister.Arguments = os.Args
	cmdRegister.Name = "register"

	cmds := models.Commands{
		Registered: make(map[string]func(s *models.State, c models.Command) error),
	}
	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)

	login, ok := cmds.Registered["login"]
	if !ok {
		os.Exit(1)
	}

	reg, ok := cmds.Registered["register"]
	if !ok {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "login":
		if err := login(&state, cmdLogin); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("User has been set.")
		}
	case "register":
		if err := reg(&state, cmdRegister); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("User registered successfully")
		}
	}
}
