package response

import "time"

type CurrentlyPlaying struct {
	context Context
	currentlyPlayingType string
	isPlaying bool
	item interface{}
	progress time.Duration
	timestamp time.Time
}

type PlayHistory struct {
	track TrackFull
	playedAt time.Time
	context Context
}

type ResumePoint struct {
	fullyPlayed bool
	resumePosition int
}
