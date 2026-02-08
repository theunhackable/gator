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

	cmdReset := models.Command{
		Arguments: os.Args,
		Name:      "reset",
	}
	cmdUsers := models.Command{
		Arguments: os.Args,
		Name:      "users",
	}

	cmds := models.Commands{
		Registered: make(map[string]func(s *models.State, c models.Command) error),
	}
	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerUsers)

	login, ok := cmds.Registered["login"]
	if !ok {
		os.Exit(1)
	}

	reg, ok := cmds.Registered["register"]
	if !ok {
		os.Exit(1)
	}

	reset, ok := cmds.Registered["reset"]
	if !ok {
		os.Exit(1)
	}

	users, ok := cmds.Registered["users"]
	if !ok {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "login":
		if err := login(&state, cmdLogin); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("User has been set.")
		}
	case "register":
		if err := reg(&state, cmdRegister); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("User registered successfully")
		}

	case "reset":
		if err := reset(&state, cmdReset); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("Db Reset done.")
		}
	case "users":
		if err := users(&state, cmdUsers); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
