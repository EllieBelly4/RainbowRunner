## interesting

// Seems to be linked to a static object

* har_traitsIwESaIwEEE

values are little endian town zone GCClassRegistry::readType() : 87A770C6

255 - 1 > 0 Loop 0 value: 0FF 77C99Ch : .text:005FC1EE ZoneClient::processConnected loop 1: 0FE 77C99

## useful

_7DFCDataNode@@6B@

??_7WorldObjectGroup@@6BDFCNode@@@

First objects are added to the world Then the resources are loaded

from .text:004D0590

if ( a2 )
{ if ( (unsigned __int8)IsKindOf<WorldObject,GCObject>(a2) )
{
(*(void (__thiscall **)(struct DFCNode *))(*(_DWORD *)a2 + 168))(a2); } else if ( (unsigned __int8)IsKindOf<
WorldEntity,DFCNode>(a2) )
{
(*(void (__thiscall **)(struct DFCNode *))(*(_DWORD *)a2 + 228))(a2); } } v3 = (struct DFCNode *)*((_DWORD *)a2 + 6);
while ( v3 )

??_7WorldObjectGroup@@6BDFCNode@@@

WorldObjectGroups seem to contain list of static objects to spawn

WorldEntity::setPosition

.text:004D3293 89 87 90 00 00 00             mov     [edi+90h], eax
.text:004D3299 8B 4E 04                      mov     ecx, [esi+4]    ; this
.text:004D329C 89 8F 94 00 00 00             mov     [edi+94h], ecx
.text:004D32A2 8B 56 08                      mov     edx, [esi+8]
.text:004D32A5 57                            push    edi
.text:004D32A6 89 97 98 00 00 00             mov     [edi+98h], edx

## back again

.text:004D0D80 ; void __thiscall World::PopulateStaticObjectPlaceholders(#360 *__hidden this, int)
.text:004D0D80 ?PopulateStaticObjectPlaceholders@World@@QAEXH@Z proc near

"placeholderstatics"
