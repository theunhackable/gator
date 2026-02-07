package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/helpers"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerRegister(s *models.State, cmd models.Command) error {
	fmt.Println("start of register")
	argLen := len(cmd.Arguments)
	expArgLen := 3
	if argLen != expArgLen {
		return helpers.ExpectedRequired(expArgLen, argLen)
	}
	username := cmd.Arguments[expArgLen-1]

	getCtx := context.Background()
	name, err := s.Db.GetUserByUsername(getCtx, username)

	if err == nil {
		return fmt.Errorf("User %s already exists.\n", name)
	}
	var newUser db.CreateUserParams

	newUser.Name = username
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	ctx := context.Background()

	user, err := s.Db.CreateUser(ctx, newUser)
	if err != nil {
		return fmt.Errorf("Error registering user: %w", err)
	}

	fmt.Printf("User %s has been registered.\n", user.Name)
	return nil
}
