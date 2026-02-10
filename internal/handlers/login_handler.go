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
	user, err := s.Db.GetUserDetailsByUsername(getCtx, username)

	if err != nil {
		fmt.Printf("User %s not found\n", user.Name)
		os.Exit(1)
	}

	if err := s.State.SetUser(user.Name); err != nil {
		return err
	}
	return nil
}
