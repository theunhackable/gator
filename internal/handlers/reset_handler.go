package handlers

import (
	"context"

	"github.com/theunhackable/gator/internal/models"
)

func HandlerReset(s *models.State, cmd models.Command) error {
	ctx := context.Background()
	if err := s.Db.ResetUserTable(ctx); err != nil {
		return err
	}

	return nil

}
