package response

type Disallows struct {
	interruptingPlayback bool
	pausing bool
	resuming bool
	seeking bool
	skippingNext bool
	skippingPrev bool
	togglingRepeatContext bool
	togglingShuffle bool
	togglingRepeatTrack bool
	transferringPlayback bool
}