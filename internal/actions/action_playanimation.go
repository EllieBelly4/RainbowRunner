package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionPlayAnimation
type ActionPlayAnimation struct {
	Unk0 byte
	/**
	--AnimationID is also offset depending on weapon desc, default is +100
	-- 0xFF = Animation ID 0
	-- 0x00 = Animation ID 0
	-- 0x02 = Animation ID v8 + 5
	-- 0x03 = Animation ID v8 + 6
	-- 0x07 = Animation ID v8 + 50
	-- 0xXX = Animation ID from_client_2 + v8
	*/
	AnimationIDSelectionType uint32
	AnimationID              uint32
	AnimationFrames          uint32
	Unk4                     uint32
}

func (d ActionPlayAnimation) OpCode() BehaviourAction {
	return BehaviourActionPlayAnimation
}

func (d ActionPlayAnimation) Init(body *byter.Byter) {
	body.WriteByte(d.Unk0)
	body.WriteUInt32(d.AnimationIDSelectionType)
	body.WriteUInt32(d.AnimationID)
	body.WriteUInt32(d.AnimationFrames)
	body.WriteUInt32(d.Unk4)
}

func NewActionPlayAnimation() *ActionPlayAnimation {
	return &ActionPlayAnimation{}
}
