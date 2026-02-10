package handlers

import (
	"context"
	"fmt"

	"github.com/theunhackable/gator/internal/helpers"
	"github.com/theunhackable/gator/internal/models"
)

func HandlerAgg(s *models.State, cmd models.Command) error {
	ctx := context.Background()
	feed, err := helpers.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
