# DFCClasses

Every useful object extends DFCClass, this can include internal types as well as types that can be sent from the server
to the client.

## Types

Each DFCClass has a type that determines how it can be used. When trying to perform an action using a class
the client will verify it's type against a mask.

### Class mask groups

These show which of the mask groups the objects match against in a 64(ish)bit mask.
The game client is 32bit and so treats the mask as 2 separate 32bit masks.

Each entry is `1<<$ARRAY_INDEX`.

[dfc_class_mask_data.json](../../Dumps/dfc_class_mask_data.json)