package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Search       []SearchResult `json:"search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"response"`
}
type SearchResult struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

func Search(apiKey, title string) (Result, error) {
	var urlValues url.Values = make(url.Values, 2)
	urlValues.Set("apiKey", apiKey)
	urlValues.Set("s", title)

	resp, err := http.Get("http://www.omdbapi.com/?" + urlValues.Encode())
	if err != nil {
		return Result{}, fmt.Errorf("failed to make request to omdb: %w", err)
	}
	defer resp.Body.Close()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("failed to decode response from omdb: %w", err)
	}

	return result, nil
}
