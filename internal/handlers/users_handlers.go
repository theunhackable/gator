package handlers

import (
	"context"
	"fmt"

	"github.com/theunhackable/gator/internal/models"
)

func HandlerUsers(s *models.State, cmd models.Command) error {

	ctx := context.Background()
	users, err := s.Db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("Error executing users command.")
	}

	currentUser := s.State.CurrentUserName

	for _, user := range users {
		if user.Name == currentUser {

			fmt.Printf("* %s (current)\n", user.Name)
		} else {

			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil

}
