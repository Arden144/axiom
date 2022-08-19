package music

import (
	"container/list"

	"github.com/disgoorg/disgolink/lavalink"
)

type Queue struct{ *list.List }

func newQueue() Queue {
	return Queue{list.New()}
}

func (q *Queue) Enqueue(tracks ...lavalink.AudioTrack) {
	for _, track := range tracks {
		q.PushFront(track)
	}
}

func (q *Queue) Dequeue() (lavalink.AudioTrack, bool) {
	e := q.Back()
	if e == nil {
		return nil, false
	}

	v := q.Remove(e)
	return v.(lavalink.AudioTrack), true
}

func (q *Queue) Remaining() lavalink.Duration {
	d := lavalink.Duration(0)
	for e := q.Front(); e != nil; e = e.Next() {
		d += e.Value.(lavalink.AudioTrack).Info().Length
	}
	return d
}

func (q *Queue) Clear() {
	q.Init()
}
