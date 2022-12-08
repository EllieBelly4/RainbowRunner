package messages

type ZoneMessage byte

const (
	ZoneMessageConnected ZoneMessage = iota
	ZoneMessageReady
	ZoneMessageDisconnected
	ZoneMessageInstanceCount ZoneMessage = 5
)
