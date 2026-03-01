package helpers

import (
	"context"
	"errors"
	"fmt"

	"github.com/theunhackable/gator/internal/models"
)

func ScrapeFeeds(s *models.State) error {
	feeds, err := s.Db.GetNextFeedToFetch(context.Background(), 1)
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		return errors.New("no feeds found.")
	}

	feed := feeds[1]

	rssFeed, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	if err := s.Db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		bufferLen := len(item.Title) + 3
		for range bufferLen {
			fmt.Print("#")
		}
		fmt.Println()
		fmt.Println("#")
		fmt.Print("#")
		fmt.Println("  ", item.Title, "  #")
		for range bufferLen {
			fmt.Print("#")
		}
		fmt.Println()
	}
	return nil
}
