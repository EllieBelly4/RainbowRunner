# Avatar position

Seems like it might be set in `Unit::setPosition`, but I cannot find where in the messages it comes from.
During EntityCreate it seems to be set to 0,0,0, assuming it must be in EntityInit?

`Unit::setPosition` retrieves 3 values for position from `Avatar as Unit` + 0x94, 0x9A, 0x9C.
`WorldEntity::readInit` reads the initial position as 3 uint32 

