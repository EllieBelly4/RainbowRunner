package messages

//go:generate stringer -type=Channel
type Channel byte

const (
	NoChannel Channel = iota
	Unk1
	Unk2
	UserChannel
	CharacterChannel
	Unk5
	ChatChannel
	ClientEntityChannel
	Unk8
	GroupChannel
	TradeChannel
	UnkB
	UnkC
	ZoneChannel
	UnkE
	PosseChannel
)
