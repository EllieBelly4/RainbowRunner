# Serialisation

Objects can be (de)serialised in several ways depending on the functions that are used to process them.

In areas where performance and bandwidth are less important the full [GCObject serialisation](#GCObjects) is used, other
areas will use a custom deserialisation method but usually still rely on the GCObject types.

## GCObject - Full Mode

GCObjects are serialised objects which deserialised into DFCNodes(?).
This is the "full mode" of serialisation where strings are used for all types, in other modes types are optionally 
strings but can be sent as hashes directly see: `GCClassRegistry::readType`.

### Object Instantiation

After the GCObject type has been determined a new instance of the class may be created using one of the `create$NativeType`
functions.

Any further processing of the objects is dependent on how it was instantiated, I think the parent usually handles any 
additional processing e.g. `Avatar::createEquipment`.

### Structure

#### Version `byte`

This is the version of the GCObject that is being sent, anything below the current version goes through a GCObject
migration function and will probably still work.

Current version in `v666` is `0x29`

#### Native (DFCNode?) object type `CString`

This is the name of the native object type that your GCObject type is going to deserialise into.

You can usually just guess the native type by looking through the GCDictionary for something that seems like the generic
type of that GCType

Example:

```
GCType      = avatar.classes.fighterfemale 
NativeType  = Avatar
```

#### ID `uint32`

ID of this GCObject, not always useful but later on this can be used as a reference for adding children or updating
specific objects/components.

#### Name `CString`

Name of this GCObject(?), I have not found a use for this, any value is accepted.

#### ChildCount `uint32`

Number of child GCObjects that are being added.

#### Children `GCObject[]`

This is a list of serialised GCObjects, GCObjects will not always allow all types of GCObjects as children.

#### GCType `CString`

This is the GCType of the object that maps to the provided NativeType.

#### Properties `GCObjectProperty[]`

##### GCObjectProperty.Name `CString`

##### GCObjectProperty.Value `uint32 | CString | byte | uint16`

The value that is required depends on the property, you can find most(all?) properties by searching for classes that
start with `Property`, the class name structure is usually `Property$NativeType$PropertyName`. They will have functions
like `getName` `asInt` `asFloat` e.g. `PropertyAvatarHair::asInt`.

#### Null `0x00`

You must end a GCObject with a null byte.