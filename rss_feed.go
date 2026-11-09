package main
import (
	"net/http"
	"fmt"
	"io"
	"context"
	"encoding/xml"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title 		string 	  `xml:"title"`
		Link 		string 	  `xml:"link"`
		Description string 	  `xml:"description"`
		Item 		[]RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title 		string `xml:"title"`
	Link 		string `xml:"link"`
	Description string `xml:"description"`
	PubDate 	string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching from feed URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body of response: %w", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("error parsing xml into feed struct: %w", err)
	}

	ch := &feed.Channel
	ch.Title = html.UnescapeString(ch.Title)
	ch.Description = html.UnescapeString(ch.Description)
	for i := range ch.Item {
		ch.Item[i].Title = html.UnescapeString(ch.Item[i].Title)
		ch.Item[i].Description = html.UnescapeString(ch.Item[i].Description)
	}

	return &feed, nil
}