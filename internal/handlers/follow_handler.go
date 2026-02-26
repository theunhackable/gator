package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/helpers"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerFollow(s *models.State, cmd models.Command) error {

	argLen := len(cmd.Arguments)
	expArgLen := 3
	if argLen != expArgLen {
		return helpers.ExpectedRequired(expArgLen, argLen)
	}
	feedUrl := cmd.Arguments[argLen-1]
	feedCtx := context.Background()
	feedDetails, err := s.Db.GetFeedDetailsByUrl(feedCtx, feedUrl)

	if err != nil {
		return err
	}

	username := s.State.CurrentUserName
	userCxt := context.Background()
	user, err := s.Db.GetUserDetailsByUsername(userCxt, username)
	if err != nil {
		return err
	}
	userId := user.ID
	feedId := feedDetails.ID

	createFeedCtx := context.Background()
	details, err := s.Db.CreateFeedFollow(createFeedCtx, db.CreateFeedFollowParams{
		UserID:    userId,
		FeedID:    feedId,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
	fmt.Printf("%s follows %s now.", details.UserName, details.FeedName)
	return nil
}
