# Invisible Player

The player is not visible in game.

[???] `Avatar::buildVisual` appears to be where the avatar should be built, `Visual::buildVisual` 
does not seem to be creating the mounted visual for Face due to 
`[MountedVisual + 18h]` being `0`.

Avatar may not be marked correctly as player controlled,
on start it sends a request for client control.
Client control appears to be managed by the UnitBehavior

ClientUnitBehavior seems important

[UnitBehaviour+156h] seems to trigger request for client control

DivorceClientFromLogicalMovement appears to be removing direct control from player
Remarry must give back control

After messing with the `WorldEntity::readInit` flags for a while the mouse zoom function broke.
Changing the flags back did not fix it