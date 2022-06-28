package music

import (
	"log"

	"github.com/disgoorg/disgolink/lavalink"
)

type eventListener struct {
	lavalink.PlayerEventAdapter
	player Player
}

func newListener(player Player) *eventListener {
	return &eventListener{player: player}
}

func (l *eventListener) OnTrackEnd(_ lavalink.Player, track lavalink.AudioTrack, endReason lavalink.AudioTrackEndReason) {
	if !endReason.MayStartNext() {
		return
	}

	if err := l.player.Next(); err != nil {
		log.Print("WARN: failed to play: ", err)
	}

}
