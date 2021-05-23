## DFCNode

### Properties

|Name|Type|
|---|---|
|Name? |CString|

## Player : Entity :? DFCNode

### Properties

Seemingly has no serialisable properties

## Avatar :? DFCNode

### Child nodes?
```
Avatar::AddNode<Modifiers>(void)
Avatar::AddNode<Manipulators>(void)
Avatar::AddNode<UnitBehavior>(ArchetypeRef<UnitBehavior> const &)
Avatar::AddNode<Skills>(ArchetypeRef<Skills> const &)
Avatar::AddNode<Equipment>(ArchetypeRef<Equipment> const &)
Avatar::AddNode<UnitContainer>(void)
Avatar::AddNode<UnitBehavior>(void)
Avatar::AddNode<Skills>(void)
Avatar::AddNode<Equipment>(void)
```
### Properties

|Name|Type|
|---|---|
|Skin|uint32|
|Face|uint32|
|FaceFeature|uint32|
|Hair|uint32|
|HairColor|uint32|
|TotalWorldTime||
|DescDeathGoldPenalty||
|DescDeathExpPenalty||
|DescSummonGoldPenalty||
|DescSummonExpPenalty||
|DescStartingCurrency||
|MetricsSaveCounter||
|Level||

## AvatarDesc :? WorldEntityDesc

## Unit :? DFCNode

### Properties

|Name|Type|
|---|---|
|Level|uint32|

## Hero : Unit

### Properties

|Name|Type|
|---|---|
|Experience|uint32|