package response

import "time"

type PlayCurrentResponse struct{}

type PlayCurrent struct {
	context              Context
	currentlyPlayingType string
	isPlaying            bool
	item                 interface{}
	progress             time.Duration
	timestamp            time.Time
}

type PlayHistoryResponse struct{}

type PlayHistory struct {
	track    TrackFull
	playedAt time.Time
	context  Context
}

type PlayHistoryPageResponse struct{}

type PlayHistoryPage struct{}

type PlayResumePointResponse struct{}

type PlayResumePoint struct {
	fullyPlayed    bool
	resumePosition int
}
