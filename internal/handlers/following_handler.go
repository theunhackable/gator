package handlers

import (
	"context"
	"fmt"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerFollowing(s *models.State, cmd models.Command, user *db.User) error {
	ctx := context.Background()
	details, err := s.Db.GetFeedFollowsForUser(ctx, user.Name)
	if err != nil {
		return err
	}

	for i, item := range details {
		fmt.Printf("%d) %s\n", i+1, item)
	}

	return nil
}
