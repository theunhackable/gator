package handlers

import (
	"context"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/helpers"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerUnfollow(s *models.State, cmd models.Command, user *db.User) error {
	argLen := len(cmd.Arguments)
	expArgLen := 3
	if argLen != expArgLen {
		return helpers.ExpectedRequired(expArgLen, argLen)
	}

	url := cmd.Arguments[2]
	ctx := context.Background()

	if err := s.Db.UnfollowFeed(ctx, db.UnfollowFeedParams{
		Url:    url,
		UserID: user.ID,
	}); err != nil {
		return err
	}
	return nil

}
