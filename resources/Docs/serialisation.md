# Serialisation

Objects can be packages using PkgFileInputStream and PackageReader.
Objects can be compressed with zlib

## Network serialisation
DFCObject/GCObject

GCClass wraps Player object?

Entity and Player are related, base/GCObject name?

### Specific types
Objects are deserialised by specialised `create` functions:

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

#### message types
```
007465DC     ; char aBasearmorclass_3[]
.rdata:007465DC     aBasearmorclass_3 db 'basearmorclasses.plate',0
.rdata:007465F3                     align 4
.rdata:007465F4     ; char aCursespalSpeed[]
.rdata:007465F4     aCursespalSpeed db 'cursespal.speedm_hineg',0
.rdata:0074660B                     align 4
.rdata:0074660C     ; char aBasearmorclass[]
.rdata:0074660C     aBasearmorclass db 'basearmorclasses.scale',0
.rdata:00746623                     align 4
.rdata:00746624     ; char aCursespalSpeed_1[]
.rdata:00746624     aCursespalSpeed_1 db 'cursespal.speedm_mdneg',0
.rdata:0074663B                     align 4
.rdata:0074663C     ; char aBasearmorclass_1[]
.rdata:0074663C     aBasearmorclass_1 db 'basearmorclasses.splint',0
.rdata:00746654     ; char aCursespalSpeed_0[]
.rdata:00746654     aCursespalSpeed_0 db 'cursespal.speedm_loneg',0
.rdata:0074666B                     align 4
.rdata:0074666C     ; char aBasearmorclass_2[]
.rdata:0074666C     aBasearmorclass_2 db 'basearmorclasses.leather',0
.rdata:00746685                     align 4
.rdata:00746688     ; char aEnhancementspa[]
.rdata:00746688     aEnhancementspa db 'enhancementspal.speedm_lo',0
.rdata:007466A2                     align 4
.rdata:007466A4     ; char aBasearmorclass_0[]
.rdata:007466A4     aBasearmorclass_0 db 'basearmorclasses.rubber',0
.rdata:007466BC     ; char aEnhancementspa_1[]
.rdata:007466BC     aEnhancementspa_1 db 'enhancementspal.speedm_md',0
.rdata:007466D6                     align 4
.rdata:007466D8     ; char aBasearmorclass_4[]
.rdata:007466D8     aBasearmorclass_4 db 'basearmorclasses.cloth',0
.rdata:007466EF                     align 10h
.rdata:007466F0     ; char aEnhancementspa_0[]
.rdata:007466F0     aEnhancementspa_0 db 'enhancementspal.speedm_hi',0
.rdata:00746740     aItemMigration8 db 'Item::Migration8',0
.rdata:00746751                     align 4
.rdata:00746754     ; char aSoulboundcount[]
.rdata:00746754     aSoulboundcount db 'SoulBoundCountdown',0
.rdata:00746767                     align 4
.rdata:00746768     ; char aNosell[]
.rdata:00746768     aNosell         db 'NoSell',0
.rdata:0074676F                     align 10h
.rdata:00746770     ; char aSoulbound[]
.rdata:00746770     aSoulbound      db 'SoulBound',0
.rdata:00738604     ; char aSkillprofessio_0[]
.rdata:00738604     aSkillprofessio_0 db 'SkillProfession',0
.rdata:00738614     ; char aRequiredprofes[]
.rdata:00738614     aRequiredprofes db 'RequiredProfession3',0
.rdata:00738628     ; char aRequiredprofes_0[]
.rdata:00738628     aRequiredprofes_0 db 'RequiredProfession2',0
.rdata:0073863C     ; char aRequiredprofes_1[]
.rdata:0073863C     aRequiredprofes_1 db 'RequiredProfession1',0
.rdata:00738650     ; char aSlotid[]
.rdata:00738650     aSlotid         db 'SlotID',0
.rdata:00738657                     align 4
.rdata:00738658     ; char aSkillprofessio[]
.rdata:00738658     aSkillprofessio db 'SkillProfessionDesc',0
```

```
.rdata:00748694     aPlayerid       db 'PlayerID',0
.rdata:0074869D                     align 10h
.rdata:007486A0     ; char aGroupid[]
.rdata:007486A0     aGroupid        db 'GroupID',0
.rdata:007486A8     ; char aItemobject[]
.rdata:007486A8     aItemobject     db 'ItemObject',0
```

#### Notes
```
05/05 11:22:12.757      37908:  FatalError: Serialization error.  Couldn't find expected property a on object with GCClassType 'player' and NativeType 'Player'
```





