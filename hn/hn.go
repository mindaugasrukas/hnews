package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// API represents the HackerNews API
	API struct {
		Endpoint string
	}

	// Item represents a HackerNews item
	Item struct {
		ID          int    `json:"id"`
		Deleted     bool   `json:"deleted"`
		Type        string `json:"type"`
		By          string `json:"by"`
		Time        int    `json:"time"`
		Text        string `json:"text"`
		Dead        bool   `json:"dead"`
		Parent      int    `json:"parent"`
		Poll        int    `json:"poll"`
		Kids        []int  `json:"kids"`
		URL         string `json:"url"`
		Score       int    `json:"score"`
		Title       string `json:"title"`
		Parts       []int  `json:"parts"`
		Descendants int    `json:"descendants"`
	}

	// PostItem composes an Item with its comments
	PostItem struct {
		*Item `json:"-"`

		Text     string     `json:"text,omitempty"`
		Comments []PostItem `json:"comments,omitempty"`
	}
)

const (
	DefaultEndpoint = "https://hacker-news.firebaseio.com/v0"
)

// NewAPI returns a new instance of the HackerNews API
func NewAPI() *API {
	return &API{
		Endpoint: DefaultEndpoint,
	}
}

// GetItem fetches a HackerNews item by ID
func (api *API) GetItem(id int) (*Item, error) {
	// send a GET request to the HackerNews API
	request := fmt.Sprintf("%s/item/%d.json", api.Endpoint, id)
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}

	// decode the response into an Item struct
	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

// GetPost fetches a HackerNews post by ID with all its comments
func (api *API) GetPost(id int) (*PostItem, error) {
	// fetch the root post
	item, err := api.GetItem(id)
	if err != nil {
		return nil, err
	}

	// create a new PostItem struct
	text := ""
	switch item.Type {
	case "story":
		text = item.Title
		if item.Text != "" {
			text += "\n" + item.Text
		}
		if item.URL != "" {
			text += "\n" + item.URL
		}
	case "comment":
		text = item.Text
	}
	post := &PostItem{Item: item, Text: text}

	// fetch all the comments
	for _, kid := range post.Kids {
		comment, err := api.GetPost(kid)
		if err != nil {
			return nil, err
		}

		post.Comments = append(post.Comments, *comment)
	}

	return post, nil
}
