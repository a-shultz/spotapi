package response

type Error struct {
	status int
	message string
}

type PlayerError struct {
	status int
	message string
	reason string // reason strings defined in object model
}