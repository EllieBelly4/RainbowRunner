# Character Selection

## Handler `CharacterManagerClient::processMessages`

### Server -> Client messages

|Message|Desc|
|---|---|
|CharacterConnected|Init connection|
|CharacterDisconnected||
|CharacterCreate|New character created, contains [GCObject](../Serialisation.md#GCObjects) with selected options|
|CharacterGetList|List of existing characters as [GCObjects](../Serialisation.md#GCObjects)|
|CharacterDelete|Delete a character|
|CharacterPlay|Start the game with selected character|
