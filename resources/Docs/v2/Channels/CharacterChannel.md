# Character Selection

**Handler**: `CharacterManagerClient::processMessages`

## Server -> Client messages

|ID|Message|Desc|
|---|---|---|
|`0x00`|Connected|Init connection|
|`0x01`|Disconnected||
|`0x02`|CharacterCreated|New character created, contains [GCObject](../Serialisation.md#GCObjects) with selected options|
|`0x03`|GotCharacter|Send list of existing characters as [GCObjects](../Serialisation.md#GCObjects)|
|`0x04`|DeleteCharacter|Delete a character|
|`0x05`|SelectCharacter|Start the game with selected character|

## Client -> Server messages

|ID|Message|Desc|
|---|---|---|


### GotCharacter `0x03`

The main deserialisation of this message happens in `Player::readObject`.