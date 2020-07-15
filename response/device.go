package response

type DevicesResponse struct{}

type Devices = []Device

type DeviceResponse struct{}

type Device struct {
	id               string
	isActive         bool
	isPrivateSession bool
	isRestricted     bool
	name             string
	deviceType       string
	volumePercent    int
}
