package search

import "errors"

type Image struct {
	Url string
}

type Album struct {
	Artists    []Artist
	Href       string
	Images     []Image
	Name       string
	TrackCount int
}

type Artist struct {
	Href string
	Name string
}

type Track struct {
	Album   Album
	Artists []Artist
	Href    string
	Name    string
}

type Result struct {
	Tracks struct {
		Href  string
		Items []Track
	}
	Next     string
	Previous string
	Total    int
}

var ErrNoResults = errors.New("no results found")

func Search(query string) (*Track, error) {
	var result Result

	params := map[string]string{
		"q":     query,
		"type":  "track",
		"limit": "1",
	}

	_, err := client.R().
		SetQueryParams(params).
		SetResult(&result).
		Get(SEARCH)

	if err != nil {
		return nil, err
	}

	if len(result.Tracks.Items) == 0 {
		return nil, ErrNoResults
	}

	return &result.Tracks.Items[0], nil
}
