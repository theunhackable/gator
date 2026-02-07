package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/theunhackable/gator/internal/helpers"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerLogin(s *models.State, cmd models.Command) error {
	argLen := len(cmd.Arguments)
	expArgLen := 3
	if argLen != expArgLen {
		return helpers.ExpectedRequired(expArgLen, argLen)
	}
	username := cmd.Arguments[expArgLen-1]
	getCtx := context.Background()
	name, err := s.Db.GetUserByUsername(getCtx, username)

	if err != nil {
		fmt.Printf("User %s not found\n", name)
		os.Exit(1)
	}

	if err := s.State.SetUser(name); err != nil {
		return err
	}
	return nil
}
