package music

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var ErrNotFound = errors.New("track not found")
var ErrInvalidUrl = errors.New("invalid url")
var ErrNotSearch = errors.New("expected search result")
var ErrNotTrack = errors.New("expected track result")
var identifierRegex = regexp.MustCompile("^.*(?:youtu.be\\/|v\\/|e\\/|u\\/\\w+\\/|embed\\/|v=)([^#\\&\\?]*).*")

type Player struct {
	disgolink.Player
	Queue
}

func (p *Player) Search(ctx context.Context, name string) ([]lavalink.Track, error) {
	n := p.Lavalink().BestNode()

	identifier := fmt.Sprint("ytsearch:", name)
	result, err := n.LoadTracks(ctx, identifier)
	if err != nil {
		return nil, err
	}

	if err, ok := result.Data.(lavalink.Exception); ok {
		return nil, err
	}

	if result.LoadType == lavalink.LoadTypeEmpty {
		return nil, ErrNotFound
	}

	tracks, ok := result.Data.(lavalink.Search)
	if !ok {
		return nil, ErrNotSearch
	}

	return tracks, nil
}

func (p *Player) ResolveUrl(ctx context.Context, url string) (lavalink.Track, error) {
	c := p.Node()

	submatches := identifierRegex.FindStringSubmatch(url)
	if len(submatches) != 2 {
		return lavalink.Track{}, ErrInvalidUrl
	}
	identifier := submatches[1]

	result, err := c.LoadTracks(ctx, identifier)
	if err != nil {
		return lavalink.Track{}, err
	}

	if err, ok := result.Data.(lavalink.Exception); ok {
		return lavalink.Track{}, err
	}

	if result.LoadType == lavalink.LoadTypeEmpty {
		return lavalink.Track{}, ErrNotFound
	}

	track, ok := result.Data.(lavalink.Track)
	if !ok {
		return lavalink.Track{}, ErrNotTrack
	}

	return track, nil
}

func (p *Player) Next() error {
	if track, ok := p.Dequeue(); ok {
		return p.Update(context.TODO(), lavalink.WithPaused(false), lavalink.WithTrack(track))
	}
	return p.Update(context.TODO(), lavalink.WithNullTrack())
}

func (p *Player) Playing() bool {
	track := p.Track()
	return track != nil && track.Info.Length != p.Position()
}
