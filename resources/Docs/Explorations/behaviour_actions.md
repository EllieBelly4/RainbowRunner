Die Action:

`.rdata:00879B78 ??_7Die@@6BDFCNode@@@ dd offset ??_GBehaviorDesc@@UAEPAXI@Z`

The Die action has an ID of 0xFF, if we can find how this is registered we can
find all IDs.

Maybe try creating a Die object from the character select or create entity so we can 
search the GCObject properties

Behaviors are registered in `.text:00513E20 ?registerClass@Registry@Action@@SAXEPAVDFCClass@@@Z`
