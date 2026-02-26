package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/helpers"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerAddFeed(s *models.State, cmd models.Command) error {

	argLen := len(cmd.Arguments)
	expArgLen := 4
	if argLen != expArgLen {
		return helpers.ExpectedRequired(expArgLen, argLen)
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

	_, feedFollowErr := s.Db.CreateFeedFollow(ctx, db.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    feed.ID,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})

	if feedFollowErr != nil {
		return feedFollowErr
	}

	fmt.Printf("feed name '%s' inserted successfully\n", feed.Name)

	return nil
}
