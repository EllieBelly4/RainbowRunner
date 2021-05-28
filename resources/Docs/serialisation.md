# Serialisation

Objects can be packages using PkgFileInputStream and PackageReader.
Objects can be compressed with zlib

## Network serialisation
DFCObject/GCObject

GCClass wraps Player object?

Entity and Player are related, base/GCObject name?

### Specific types
Objects are deserialised by specialised `create` functions:

### GCClassRegistry

This registry contains definitions for all GCObjects, a dictionary of all possible types is stored as a PKG file inside game.pkg.
[gcdictionary.dict](../Dumps/010/game_pkg_gcdictionary.dict_uncompressed_body)

### GCParser

#### Potential keywords

* extends
* synchronized
* function
* var
* static
* transient
* state
* event

### GCObject

GCObjects are serialisable objects and read from serialised strings containing property names/types and values.

#### Properties

It seems there are "PropertyX" classes for some serialisable types.

e.g. PropertyAvatarHair

They implement a `getName` and `setValue` method.

#### GCNativeProperty

The name property found on the player seems to be allowed because this property exists.

##### GCNativeProperty Parsers
StringProperty::read

##### Properties
###### NameGCObjectName `StringProperty`
Calls DFCNode::setName `.text:005C3F40`

#### GCMethods

Potentially these are used for serialisation?

#### Message structure

```
Start of message
[01]             Unk
["Player" 00]    DFCObject Class name
[11 12 13 14]    Unk
[EE EE]          Unk
[01 00 00 00]    Unk Length?

Property?
[01]             Unk
["int" 00]       Property Type (lowercase only, '.' has special meaning as first char)
["name" 00]      Property Name

// Same type as first ObjectTypeByte, for object parent?
body.WriteByte(0x00)
```

#### Serialisable Objects

```
GCObject
WorldObject
WorldObjectDesc
Visual
Entity
EntityDesc
CurveTable
Hero extends Avatar
HeroDesc
Unit
UnitDesc
EntityComponent
EntityComponentDesc
UnitBehavior
UnitBehaviorDesc
LevelTitle
Skill
SkillDesc
SubEntity
AreaTrigger
WorldEntityGenerator
Service
Item
ItemDesc
ItemModifier
ItemModifierDesc
Weapon
WeaponDesc
ItemGenerator
SingleItemGenerator
QuestObjective
RoomNode
GCObjectIterator
```
### DiskPkgFileInfo

Size: 0x20

```
mov     eax, ecx
and     byte ptr [eax+14h], 0FEh
and     dword ptr [eax+18h], 80000000h
xor     ecx, ecx
mov     dword ptr [eax], 0FFFFFFFFh
mov     [eax+4], ecx
mov     [eax+8], ecx
mov     dword ptr [eax+0Ch], 0FFFFFFFFh
mov     [eax+10h], ecx
mov     [eax+24h], ecx
mov     [eax+28h], ecx
mov     [eax+20h], ecx
mov     [eax+1Ch], ecx
```