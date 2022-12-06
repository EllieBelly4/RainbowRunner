package messages

import (
	"errors"
	"fmt"
)

type UndeliveredMessageNotificationReasonString byte

const (
	UndeliveredMessageNotificationReasonNoReason UndeliveredMessageNotificationReasonString = iota
	UndeliveredMessageNotificationReasonUnknownTargetDomain
	UndeliveredMessageNotificationReasonUnableToFindSender
	UndeliveredMessageNotificationReasonUnauthorizedBroadcast
	UndeliveredMessageNotificationReasonNoTargets
	UndeliveredMessageNotificationReasonUnableToFindTargetSessionId
	UndeliveredMessageNotificationReasonTargetNotLoggedIn
	UndeliveredMessageNotificationReasonTargetIsIgnoringSender
	UndeliveredMessageNotificationReasonTargetNotFound
	UndeliveredMessageNotificationReasonChatSystemUnavailable
)

// Only to be used when sending responses back to the client
type MessageChannelSource byte

const (
	MessageChannelSourceWorld MessageChannelSource = iota + 2
	MessageChannelSourceZone
	MessageChannelSourceGroup
	MessageChannelSourceTell
	MessageChannelSourceTell2 // Part of Tell? Sends back "To {NAME}" e.g. "To Testy" in pink
	MessageChannelSourceTell3 // Part of Tell, sends back "Tell> Testy"
	MessageChannelSourceUnk1
	MessageChannelSourceUnk2
	MessageChannelSourceDeliverFailure
	MessageChannelSourceMarket
	MessageChannelSourceNoob
	MessageChannelSourceGlobalAnnouncement
)

type ClientMessageChannelSource byte

const (
	ClientMessageChannelSourceWorld ClientMessageChannelSource = iota + 1
	ClientMessageChannelSourceZone
	ClientMessageChannelSourceGroup
	ClientMessageChannelSourceTell
	ClientMessageChannelSourceMarket
	ClientMessageChannelSourceNoob
	ClientMessageChannelSourcePVP
)

var clientMessageChannelSourceMap = map[ClientMessageChannelSource]MessageChannelSource{
	ClientMessageChannelSourceWorld:  MessageChannelSourceWorld,
	ClientMessageChannelSourceZone:   MessageChannelSourceZone,
	ClientMessageChannelSourceGroup:  MessageChannelSourceGroup,
	ClientMessageChannelSourceMarket: MessageChannelSourceMarket,
	ClientMessageChannelSourceNoob:   MessageChannelSourceNoob,
	//ClientMessageChannelSourcePVP: MessageChannelSourcePVP,
}

func (s ClientMessageChannelSource) ToMessageChannelSource() (MessageChannelSource, error) {
	source, ok := clientMessageChannelSourceMap[s]

	if !ok {
		return MessageChannelSourceNoob,
			errors.New(
				fmt.Sprintf("could not find message channel source for client message source %d", byte(s)),
			)
	}

	return source, nil
}
