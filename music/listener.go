package music

import (
	"log"

	"github.com/disgoorg/disgolink/lavalink"
)

type EventListener struct {
	lavalink.PlayerEventAdapter
	player Player
}

func NewListener(player Player) *EventListener {
	return &EventListener{player: player}
}

func (l *EventListener) OnTrackEnd(_ lavalink.Player, track lavalink.AudioTrack, endReason lavalink.AudioTrackEndReason) {
	if !endReason.MayStartNext() {
		return
	}

	if err := l.player.Next(); err != nil {
		log.Print("WARN: failed to play: ", err)
	}

}
