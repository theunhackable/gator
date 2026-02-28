package middleware

import (
	"context"

	"github.com/theunhackable/gator/internal/db"
	"github.com/theunhackable/gator/internal/models"
)

func MiddlewareLoggedIn(handler func(s *models.State, cmd models.Command, user *db.User) error) func(*models.State, models.Command) error {
	return func(s *models.State, cmd models.Command) error {
		username := s.State.CurrentUserName
		user, err := s.Db.GetUserDetailsByUsername(context.Background(), username)
		if err != nil {
			return err
		}
		return handler(s, cmd, &user)
	}
}
