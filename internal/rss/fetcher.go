package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}


// test url = https://hnrss.org/frontpage

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	req, err := http.NewRequestWithContext(ctx, "GET",feedURL,nil)
	if err != nil{
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	res, err := client.Do(req)
	
	if err != nil{
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("request not succesful, status code: %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	var delivery RSSFeed
	err =xml.Unmarshal(data, &delivery)
	if err != nil{
		return nil, err
	}
	delivery.Channel.Title = html.UnescapeString(delivery.Channel.Title )
	delivery.Channel.Description = html.UnescapeString(delivery.Channel.Description )
	for i := range delivery.Channel.Item{
		delivery.Channel.Item[i].Title = html.UnescapeString(delivery.Channel.Item[i].Title)
    	delivery.Channel.Item[i].Description = html.UnescapeString(delivery.Channel.Item[i].Description)
	}
	return &delivery, nil
}
