# Channels

The main form of communication after the AuthGateway is through a concept known as Channels. The channels are split up
to handle messages for different areas of the game e.g. character controls, entity management, user/friend actions.

## Available channels

1. NoChannel `0x00`
1. [UserChannel](UserChannel.md) `0x03`
1. [CharacterChannel](CharacterChannel.md) `0x04`
1. [ChatChannel](ChatChannel.md) `0x06`
1. [ClientEntityChannel](ClientEntityChannel.md) `0x07`
1. [GroupChannel](GroupChannel.md) `0x09`
1. [TradeChannel](TradeChannel.md) `0x0A`
1. [ZoneChannel](ZoneChannel.md) `0x0D`
1. [PosseChannel](PosseChannel.md) `0x0F`