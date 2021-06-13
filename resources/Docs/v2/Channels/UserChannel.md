# User Channel

This channel handles messages related to the user, this includes friend list actions and block lists and probably more.

**Handler**: `UserManagerClient::processMessages`

## Server -> Client messages

|ID|Message|Desc|
|---|---|---|
|`0x00`|ConnectedMessage| |
|`0x01`|RostersMessage| |
|`0x02`|AddContactMessage| |
|`0x03`|RemoveContactMessage| |
|`0x04`|AddIgnoreMessage| |
|`0x05`|RemoveIgnoreMessage| |
|`0x06`|RosterNotifyMessage| |
|`0x07`|RosterPropertyChangedMessage| |
|`0x08`|RostersWritableMessage| |
|`0x09`|UsersMessage| |
|`0x0A`|UserListEventMessage| |
|`0x0B`|FriendsPublicityMessage| |

## Client -> Server messages

|ID|Message|Desc|
|---|---|---|