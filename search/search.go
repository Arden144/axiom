package search

import (
	"context"
	"errors"
)

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

type Playlist struct {
	Tracks struct {
		Href  string
		Items []struct {
			Track Track
		}
	}
	Images       []Image
	Href         string
	Name         string
	ExternalURLs struct {
		Spotify string
	} `json:"external_urls"`
}

type Result struct {
	Tracks struct {
		Href  string
		Items []Track
	}
	Playlists struct {
		Href  string
		Items []Playlist
	}
	Next     string
	Previous string
	Total    int
}

var ErrNoResults = errors.New("no results found")

func Search(ctx context.Context, query string) (*Track, error) {
	var result Result

	params := map[string]string{
		"q":     query,
		"type":  "track",
		"limit": "1",
	}

	_, err := client.R().
		SetContext(ctx).
		SetQueryParams(params).
		SetSuccessResult(&result).
		Get(SEARCH)

	if err != nil {
		return nil, err
	}

	if len(result.Tracks.Items) == 0 {
		return nil, ErrNoResults
	}

	return &result.Tracks.Items[0], nil
}

func SearchPlaylist(ctx context.Context, query string) (*Playlist, error) {
	var result Result

	params := map[string]string{
		"q":     query,
		"type":  "playlist",
		"limit": "1",
	}

	_, err := client.R().
		SetContext(ctx).
		SetQueryParams(params).
		SetSuccessResult(&result).
		Get(SEARCH)

	if err != nil {
		return nil, err
	}

	if len(result.Playlists.Items) == 0 {
		return nil, ErrNoResults
	}

	var playlist Playlist

	_, err = client.R().
		SetContext(ctx).
		SetSuccessResult(&playlist).
		Get(result.Playlists.Items[0].Href)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}
