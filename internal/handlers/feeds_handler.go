package handlers

import (
	"context"
	"fmt"

	"github.com/theunhackable/gator/internal/models"
)

func HandlerFeeds(s *models.State, cmd models.Command) error {
	ctx := context.Background()
	details, err := s.Db.GetFeedDetails(ctx)

	if err != nil {
		return err
	}

	fmt.Println("Username \t Feed Name \t Feed URL")
	for _, detail := range details {
		fmt.Printf("%s \t %s \t %s\n", detail.Username, detail.FeedName, detail.Url)
	}
	return nil
}
