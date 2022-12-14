// Code generated by "stringer -type=Channel"; DO NOT EDIT.

package messages

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoChannel-0]
	_ = x[Unk1-1]
	_ = x[Unk2-2]
	_ = x[UserChannel-3]
	_ = x[CharacterChannel-4]
	_ = x[Unk5-5]
	_ = x[ChatChannel-6]
	_ = x[ClientEntityChannel-7]
	_ = x[Unk8-8]
	_ = x[GroupChannel-9]
	_ = x[TradeChannel-10]
	_ = x[UnkB-11]
	_ = x[UnkC-12]
	_ = x[ZoneChannel-13]
	_ = x[UnkE-14]
	_ = x[PosseChannel-15]
}

const _Channel_name = "NoChannelUnk1Unk2UserChannelCharacterChannelUnk5ChatChannelClientEntityChannelUnk8GroupChannelTradeChannelUnkBUnkCZoneChannelUnkEPosseChannel"

var _Channel_index = [...]uint8{0, 9, 13, 17, 28, 44, 48, 59, 78, 82, 94, 106, 110, 114, 125, 129, 141}

func (i Channel) String() string {
	if i >= Channel(len(_Channel_index)-1) {
		return "Channel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Channel_name[_Channel_index[i]:_Channel_index[i+1]]
}
