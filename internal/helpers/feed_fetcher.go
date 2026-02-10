package helpers

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/theunhackable/gator/internal/models"
)

func FetchFeed(ctx context.Context, feedURL string) (*models.RSSFeed, error) {

	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("%s\n", res.Status)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	var feed models.RSSFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	return &feed, nil
}
