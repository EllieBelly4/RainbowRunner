# Messages

There are various message formats used some are compressed, and some are not. The message type used mostly seems to
depend on whether the messages are part of the AuthGateway or if they are channel specific or not related to a channel.

Multiple messages can appear in the same packet end to end so thankfully most messages have clearly defined lengths.

## Compressed Channel Message `0x0A`

This is the format I use mostly for sending channel messages to the client, the client will send messages in this format
only for messages which do not appear to be sent to a specific channel(?).

### Structure

#### Packet Type `byte = 0x0A`

#### Unk (ClientID? DestID?) `uint24`

#### Compressed Message Length `uint32`

This is the length of the compressed message body + 7. The +7 I believe is to account for the next 7 bytes of the
packet.

#### Unk `byte`

Currently using `0x01` or `0x00`.

#### Message Type(?) `byte`

Currently using `0x0f`, `0x02` or `0x03`.

#### Unk `byte`

Currently using `0x00` for all.

#### Uncompressed message body length `uint32`

Length of the message body before compression.

#### Compressed message body `byte[]`

## Compressed Channel Message `0x0E`

This is the message format most often sent by the client for channel specific messages.

### Structure

#### Packet Type `byte = 0x0E`

#### Client ID(?) `uint24`

#### Entire Message Length `uint32`

To get the compressed body length subtract 12 from this value.

#### Unk `uint24`

This value was previously sent by the server and seems to be replayed by the client.

#### Unk `uint16`

#### Unk `uint24`

#### Uncompressed Message Body Length `uint32`

#### Compressed message body `byte[]`

