package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerAddFeed(s *models.State, cmd models.Command) error {
	if len(cmd.Arguments) < 4 {
		return fmt.Errorf("not enough arguments.")
	}

	name := cmd.Arguments[2]
	url := cmd.Arguments[3]

	username := s.State.CurrentUserName
	userCtx := context.Background()
	user, err := s.Db.GetUserDetailsByUsername(userCtx, username)
	if err != nil {
		return err
	}

	newFeed := db.CreateFeedParams{
		UserID:    user.ID,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       url,
	}

	ctx := context.Background()
	feed, err := s.Db.CreateFeed(ctx, newFeed)

	if err != nil {
		return err
	}

	fmt.Printf("feed name '%s' inserted successfully\n", feed.Name)

	return nil
}
