# Equipping items and visual display

1. World::update
1. ClientEntityManager::update
1. ClientEntityManager::updateEntities
1. Unit::update
1. Equipment::update


*::processRequest appears to be where client messages are generated?

Adding an item to the equipment list seems to be enough to associate it with a slot.
The slot it is associated with is the one defined in the definition text files.

Weapon::attach() is called when adding a weapon.
Weapon::buildVisual is called.
It appears that EntityObject::addVisual is useful.

`.text:00597A58 004 cmp     dword ptr [edi+68h], 0Bh`

Checking if slot type is 0x0B which is shield slot. for special case in Weapon::buildVisual.

## Adding and removing equipment directly

If you click on some equipment the client sends a message with action 0x29 to unequip.
Server can send back an update message for the equipment with 0x29 to unequip that item, or 0x28 to equip the item maybe.


## Inventory

Clicking on items in inventory sends a request via UnitContainer::getItem



## Hierarchy

```
Avatar
    Equipment
        EquipmentDesc
            EquipmentSlot
```


## Properties
GCObjectReader::setProperty
DFCClass::getPropertyByID

## What is an entity component?

This is the test to determin if a DFCClass is an EntityComponent

Testing on 0x7F

eax = another DFCClass

1. edx >> 5
2. eax + edx*4+0x38
3. edx << 5
4. edx & 0x1F

```
.text:004FBBCA 004 mov     ecx, [esi+DFCClass.ClassType]
.text:004FBBD0 004 mov     edx, ecx
.text:004FBBD2 004 shr     edx, 5
.text:004FBBD5 004 mov     eax, [eax+edx*4+38h]
.text:004FBBD9 004 mov     esi, edx
.text:004FBBDB 004 shl     esi, 5
.text:004FBBDE 004 sub     ecx, esi
.text:004FBBE0 004 mov     esi, 1
.text:004FBBE5 004 shl     esi, cl
.text:004FBBE7 004 and     eax, esi
.text:004FBBE9 004 cmp     eax, esi
.text:004FBBEB 004 jnz     short loc_4FBBF1
```