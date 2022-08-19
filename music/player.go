package music

import (
	"context"
	"errors"
	"fmt"

	"github.com/disgoorg/disgolink/lavalink"
)

var ErrNotFound = errors.New("track not found")

type Player struct {
	lavalink.Player
	Queue
}

func (p *Player) Search(ctx context.Context, name string) ([]lavalink.AudioTrack, error) {
	l := p.Node().Lavalink()
	c := l.BestRestClient()

	identifier := fmt.Sprint("ytsearch:", name)
	result, err := c.LoadItem(ctx, identifier)
	if err != nil {
		return nil, err
	}

	tracks := make([]lavalink.AudioTrack, len(result.Tracks))
	for i, track := range result.Tracks {
		tracks[i], err = l.DecodeTrack(track.Track)
		if err != nil {
			return nil, err
		}
	}

	switch result.LoadType {
	case lavalink.LoadTypeNoMatches:
		return nil, ErrNotFound

	case lavalink.LoadTypeLoadFailed:
		return nil, *result.Exception
	}

	return tracks, nil
}

func (p *Player) Next() error {
	if track, ok := p.Dequeue(); ok {
		if err := p.Pause(false); err != nil {
			return err
		}
		return p.Play(track)
	}
	return p.Stop()
}

func (p *Player) Playing() bool {
	track := p.PlayingTrack()
	return track != nil && track.Info().Length != p.Position()
}
