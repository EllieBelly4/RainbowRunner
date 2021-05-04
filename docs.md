### Notes

#### Auth Messages

linAC [server] -> [client] ?

linAQ [client] -> [server] ?

Receive Data Raw: 0A4A19B0

Min Message length of 5 Length is Little Endian

Message structure:

| Len | Data|
|---|---|
|06 00| 01 01 01 01 01 01| 

#### Encryption

Login info before encryption:

```
0019B350  7F 00 00 01 00 00 00 00  00 00 00 00 68 65 6C 6C  ............hell
0019B360  6F 00 00 00 00 00 00 00  00 00 31 00 00 00 00 00  o.........1.....
0019B370  00 00 00 00 00 00 00 00  00 00 19 00 56 FD BA 11  ............Výº.
```

`0019B264` Keys?

https://paginas.fe.up.pt/~ei10109/ca/des.html

Encrypted data passed in with length 0x1E Limited to 0x18 due to 8 byte block size

ESI input plain text 0019B35C

Message buffer for serialisation (works for server list): 0A4DA230

DFCMessage write buffer starts: 0A4F79F8

#### How do the clients send the data?

Then send via the message queue.

MessageQueue::enQueue

#### How do the clients read the data?

They each seem to poll the TChannelManager::hasMoreMessages for messages with a different "channel?" value

```
ZoneClient::processMessages
ChatClient::update 
UserManagerClient::update
CharacterManagerClient::update
GroupClient::update -> GroupClient::updateChannel
```

Sending messages before being "connected" causes the DFCMessageClient::updateReceivedMessages to log "Not connected,
Ignoring message"

DFCMessageClient::updateReceivedMessages seems to be base class for handling all channel messages

#### Game connection flow

1. [Client] Sends connection request

```
                                             Channel?
            [Pkt Type] [Unk    ] [Pkt Len ] [Unk]  [Unk       ] [One time key]
   00000000  02          ce 56 0a  09 00 00   0a     00 00 00 01  ef be ad de    |..V.............|
            [Null]
   00000010  00     |.|
```

2. [Server] Sends login success message

```
           [Pkt Type] [Unk    ] [Pkt Len ]  [Unk] [Msg Type?]
   00000000 02         00 e6 11  01 00 00    00    03         |.........|
```

3. [Server] Sends client connected message with client ID Message below is compressed with zlib

```
Body: b4 b3 b2 00

0a 33 32 31 17 00 00 00  00 03 00 04 00 00 00 78
9c 62 61 66 62 04 04 00  00 ff ff 00 22 00 0b   
```

4. [Client] Possibly heartbeat

```
(GameServer)Received: 
00000000  02 f7 68 0a 03 00 00 0a  00 02 00                 |..h........|
```

##### Need to assign an address (This was happening due to gameserver dying, probably not part of the flow)

When forcing the processConnected function call we get this message:
"Message sent before acquiring an address"

##### DFCMessageClientProcessConnected()

How do we get there?

Stack:

* DFCMessageClient::processConnected() 005d6fe0
* DFCMessageClient::updateReceivedMessage 005d6dd0

```
// DFCMessage + 0x56 == 3
if (*(short *)((int)local_14 + 0x56) == 3)
``` 

CharacterManager::Connect 005ab5e0

##### Channel based messages

DFCMessageServer::getClientChannelByID

#### Message Types

##### Server
|ID|Name|Description|
|---|---|---|
|0x01|Unk|Causes disconnections in all cases|
|0x02|Channel Message?|Does not work until game is in "connected" state|
|0x06|Unk||
|0x10|Auth message|Used prior to channel messages|
|0x0a|Some compressed message|Unk|

##### Client
|ID|Name|Description|
|---|---|---|
|0x02|Client messages|connect/disconnect/assign ID/heartbeat|
|0x06|Client messages?|First message for character selection uses this|

##### Message structure

Some messages are compressed with zlib and have varying structures

###### 0x06 - Unk
Seems to be client only, send 0x0a in response from server

Basic Header: 8 bytes 

```
debug430:0A4377D0                     dw 0F0B4h; destination_channel
debug430:0A4377D0                     dw 0B2B3h; unk_23

Client
[Type] [Src or Dst] [MsgLength] [Chan?] [Dst] [Unk  ] [Unk  ] [Unk] [ Unk ] [Body ]
 06     55 65 0A     0A 00 00    0A      B4    B3 B2   01 00   01    00 00   03 00
 
Server
[Type] [Src or Dst] [MsgLength] [Chan?] [Dest] [Unk  ] [Unk  ] [Unk] [Unk  ] 
 06     b4 b3 b2     0c 00 00    01      cd     b4 b3   dc ac   0d    f0 b0
 
 // Anything after the final f0 b0 will be parsed again from 0d                                   
 06     b4 b3 b2     0c 00 00    01      cd     b4 b3   dc ac   0d    f0 b0     b0 60 50 80 70
 
 // This message causes the decompression process even though it's not compressed
 06 fb 68 0a 07 00 00 0a  b4 b3 b2 01 00 b4 00 00                                  
 06 fb 68 0a 07 00 00 0a  b4 b3 b2 01 00 01 00 00                                  
```

###### (Compressed) 0x0A - Channel target message?

```
[Type] [Unk     ] [Comp Body Len + 7] [Chan] [Msg Type] [Uncomp Body Len] [Comp Data]
 0a     31 32 33   14 00 00 00         00     03         04 00 00 00       ...
```

###### (Compressed) 0x08 - Direct message?

```
[Type] [Unk     ] [Comp Body Len + 7] [Uncomp Body Len] [Comp Data]
 0a     31 32 33   14 00 00 00         04 00 00 00       ...
```

```
body.WriteByte(0xFF)
body.WriteUInt24(0xFFFFFF)
body.WriteUInt24(0xFFFFFF)
WriteMessage(0x0a, 0x313233, 0xFF, conn, body)
DFCSocketChannel::updateReceiveQueue Closing socket 127.0.0.1:2603 because uncompressed packet size 4294967295 exceeds max 1048576
```

#### Game Connection States

The game manages connections in terms of states, they start at 0 and depending on server responses generally move up.

##### GatewayClient

|State|Description|Method|
|---|---|---|
|0|Disconnected||
|1|Ready/waiting for auth|GatewayClient::LogIn|
|3|Connected to gameserver|GatewayClient::ConnectToGameServer|
|4|Connection to gameserver authorised|GatewayClient::UpdateAuthorize|
|5|Unk Error|GatewayClient::update|

##### DRAuthClient

|State|Description|Method|
|---|---|---|
|0|Disconnected||
|1|Ready/waiting for auth|DRAuthClient::Login|
|2|Connected to auth server|DRAuthClient::OnConnected|
|3|Authenticated|DRAuthClient::RecvLoginOk|
|4|Received server list|DRAuthClient::RecvServerListEx|
|5|Selected game server, waiting for server|DRAuthClient::SelectGameServer|
|6|Unk, assuming some fail message from server|Unk|
|7|Unk, assuming some fail message from server|Unk|
|8|Authorised for play|DRAuthClient::RecvPlayOk|

### Entities

Entity messages are handled by the ServerEntityManager There are 3 types of entity messages:

1. Update
2. Init
3. Remove

### Logging

Basic state logs are always sent to Logs\DungeonRunners.log Trace logs are sent to Logs\Trace.log but seem to be
disabled by default

### Program arguments

Maybe:
/authserver= /version

### Interesting string constants

```
.rdata:00707080 00000006 C CHUNK
.rdata:0070719C 00000006 C FATAL
.rdata:007071B8 00000006 C ERROR
.rdata:007071D0
00000006 C DEBUG
```

### ZoneClient

```
.rdata:0075C334 00000024 C %s(): ERROR: Unexpected message %d.
.rdata:0075C2E0 00000053 C %s(): ERROR: Ignored message from wrong ZoneServer(%s). Expected message from(%s)
.rdata:0075C284 00000039 C %s() InvalidZoneName: %s. Received from(%s). Status: %d
.rdata:0075C358 00000042 C %s(): Received message %d from ZoneServer(%s). Current status: %d
.rdata:0075C2E0 00000053 C %s(): ERROR: Ignored message from
wrong ZoneServer(%s). Expected message from(%s)
```

#### AuthServer

```
.rdata:0075B2AC 00000027 C %s Received unexpected data. State %d.
.rdata:0075D698 00000028 C %s received unexpected message type: %d
.rdata:0075B484 00000029 C %s Unexpected packet received. State %d.
.rdata:0075B56C 0000002B C %s Unexpected packet. Reason %u State %d.
.rdata:0075B020 00000032 C %s Received unexpected data. State %d. Reason %d.
.rdata:0075B390 00000033 C %s Unexpected packet received. Reason %u State %d.
```

### MessageServer

```
.rdata:0075F5F4 0000003E C %s Unexpected admin message %d from client with Address(1.%d)
.rdata:0075F6A0 0000005D C %s Message(0x%p) Channel(0x%p): ERROR Unexpected destination server id. Source(%s) Dest0(%s)
```

### NetSession reliable packets

```
.rdata:0070F168 00000045 C Received a packet out of sequence (%lu, expected %lu) - packet lost!
```

### ClientEntityManager

```
.rdata:007590F0 00000047 C ClientEntityManager::processMessage ERROR: Unexpected message size: %d
```

### STOChunkFileReader

```
.rdata:00706D5C 00000029 C Expected bool but got \"%s\" in chunk \"%s\"
.rdata:00706D88 00000029 C Expected byte but got\"%s\" in chunk \"%s\"
.rdata:00706DB4 00000029 C Expected char but got \"%s\" in chunk \"%s\"
.rdata:00706E0C 0000002A C Expected int16 but got \"%s\" in chunk \"%s\"
.rdata:00706E64 0000002A C Expected int32 but got \"%s\" in chunk \"%s\"
.rdata:00706EBC 0000002A C Expected int64 but got \"%s\" in chunk \"%s\"
.rdata:00706EE8 0000002C C Expected float32 but got \"%s\" in chunk \"%s\"
.rdata:00706F18 0000002C C Expected float64 but got \"%s\" in chunk \"%s\"
```