package handlers

import (
	"fmt"

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
	err := s.State.SetUser(username)

	if err != nil {
		return err
	}
	fmt.Println("User has been set.")
	return nil
}
