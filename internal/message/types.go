package message

type AuthServerMessage int
type AuthClientMessage int

// Server are linAC*
//go:generate stringer -type=AuthServerMessage
const (
	AuthServerProtocolVerPacket               AuthServerMessage = 0x00
	AuthServerLoginFailPacket                 AuthServerMessage = 0x01
	AuthServerBlockedAccountPacket            AuthServerMessage = 0x02
	AuthServerLoginOkPacket                   AuthServerMessage = 0x03
	AuthServerSendServerListExPacket          AuthServerMessage = 0x04
	AuthServerSendServerFailPacket            AuthServerMessage = 0x05
	AuthServerPlayFailPacket                  AuthServerMessage = 0x06
	AuthServerPlayOkPacket                    AuthServerMessage = 0x07
	AuthServerAccountKickedPacket             AuthServerMessage = 0x08
	AuthServerBlockedAccountWithMessagePacket AuthServerMessage = 0x09
	AuthServerCSCCheckPacket                  AuthServerMessage = 0x0A
	AuthServerQueueSizePacket                 AuthServerMessage = 0x0B
	AuthServerHandoffToQueuePacket            AuthServerMessage = 0x0C
	AuthServerPositionInQueuePacket           AuthServerMessage = 0x0D
	AuthServerHandoffToGamePacket             AuthServerMessage = 0x0E
)

// Client are linAQ*
//go:generate stringer -type=AuthClientMessage
const (
	AuthClientLoginPacket          AuthClientMessage = 0x00
	AuthClientAboutToPlayPacket    AuthClientMessage = 0x02
	AuthClientLogoutPacket         AuthClientMessage = 0x03
	AuthClientServerListExtPacket  AuthClientMessage = 0x05
	AuthClientSCCheckPacket        AuthClientMessage = 0x06
	AuthClientConnectToQueuePacket AuthClientMessage = 0x07
)
