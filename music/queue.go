package music

import (
	"container/list"

	"github.com/disgoorg/disgolink/v3/lavalink"
)

type Queue struct{ *list.List }

func NewQueue() Queue {
	return Queue{list.New()}
}

func (q *Queue) Enqueue(tracks ...lavalink.Track) {
	for _, track := range tracks {
		q.PushFront(track)
	}
}

func (q *Queue) Dequeue() (lavalink.Track, bool) {
	e := q.Back()
	if e == nil {
		return lavalink.Track{}, false
	}

	v := q.Remove(e)
	return v.(lavalink.Track), true
}

func (q *Queue) Remaining() lavalink.Duration {
	d := lavalink.Duration(0)
	for e := q.Front(); e != nil; e = e.Next() {
		d += e.Value.(lavalink.Track).Info.Length
	}
	return d
}

func (q *Queue) Clear() {
	q.Init()
}
