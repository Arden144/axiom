package music

import (
	"context"
	"errors"
	"fmt"

	"github.com/disgoorg/disgolink/disgolink"
	"github.com/disgoorg/disgolink/lavalink"
)

var ErrNotFound = errors.New("track not found")

type Player struct {
	link   disgolink.Link
	player lavalink.Player
	queue  Queue
}

func (p *Player) Play(ctx context.Context, name string) (*lavalink.AudioTrackInfo, error) {
	c := p.link.BestRestClient()

	identifier := fmt.Sprint("ytsearch:", name)
	result, err := c.LoadItem(ctx, identifier)
	if err != nil {
		return nil, err
	}

	tracks := make([]lavalink.AudioTrack, len(result.Tracks))
	for i, track := range result.Tracks {
		tracks[i], err = p.link.DecodeTrack(track.Track)
		if err != nil {
			return nil, err
		}
	}

	switch result.LoadType {
	case lavalink.LoadTypeTrackLoaded:
		// tracks[0]

	case lavalink.LoadTypePlaylistLoaded:
		// NewAudioPlaylist(result.PlaylistInfo.Name, result.PlaylistInfo.SelectedTrack, tracks)

	case lavalink.LoadTypeSearchResult:
		// tracks

	case lavalink.LoadTypeNoMatches:
		return nil, ErrNotFound

	case lavalink.LoadTypeLoadFailed:
		return nil, *result.Exception
	}

	track := tracks[0]
	info := track.Info()
	return &info, p.player.Play(tracks[0])
}

func (p *Player) AddToQueue(name string) {

}

func (p *Player) Pause() {

}

func (p *Player) Resume() {

}
