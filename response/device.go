package response

type Device struct {
	id string
	isActive bool
	isPrivateSession bool
	isRestricted bool
	name string
	deviceType string
	volumePercent int
}