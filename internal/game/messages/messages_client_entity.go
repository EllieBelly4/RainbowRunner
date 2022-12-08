package messages

type ClientEntityMessage byte

const (
	ClientEntityUnk0 ClientEntityMessage = iota
	ClientEntityUnk1
	ClientEntityUnk2
	ClientEntityUnk3
	ClientRequestRespawn
	ClientEntityUnk5
	ClientEntityUnk6
	ClientEntityUnk7
	ClientEntityUnk8
	ClientEntityUnk9
	ClientEntityComponentUpdate = 0x34
	ClientEntityMovement        = 0x35
)
