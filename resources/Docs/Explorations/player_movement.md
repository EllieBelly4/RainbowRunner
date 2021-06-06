Player cannot move.
Clicking and WASD do not send messages to the server.

When there is a synch error the ZoneClient fails saying "ZoneClient::UpdateWorld Failed to enter world. ZoneClient Status: X Last EntityManager error: Y".
Maybe this is indicating we have not fully completed the world entry procedure.

I think the above message may just be misleading, we appear to be at the highest state number in ZoneClient.


// This is unrelated, and happens when i send back the wrong response to a respawn after desynch
I sent some random data and this happened "ZoneClient::processMessages(): Received message 8 from ZoneServer(1.15). Current State: 117"


When trying to execute any actions against the player it fails due to the following, this may be related.
ESI is the UnitBehavior class
`.text:005152DC test    byte ptr [esi+7Ch], 1 ; This is not set to true and causes everything to fail`

Looking around areas which call UnitBehavior::getClientBehavior, looking for usages of ClientUnitBehavior::ResumeClientMovement

`.text:00518990` `ClientUnitBehavior.potential_desired_input_blocker` may be because the update number being sent is incorrect
`ClientEntityManager::getUpdateNumber` 

Calling moveTo or anything that affects movement when there is no ground above 0,0,0 causes you to snap to 0,0,0.
I think there is a secondary coordinate for the pathfinder that is still set to 0,0

WarpTo xyz and MoveTo xy coordinates are on the same scale... maybe.

MoveTo positions appear to be absolute target coordinates, 
the server must send these coordinates repeatedly to continue movement.
The client can move for about 0.5 seconds before waiting for server messages,
maybe this is related to the pathmanager budget?

Potential smallest grid value 0x400?

// Received player move messages may be relative, but server is expected to send absolute positions.
// The relative positions received too high, but dividing by 20 like the exp does seem to give a reasonable looking offset.
Received player move messages seem to be absolute but when it is expecting updates and how is unknown

The received positions seem to contain distance along path, which increments
by 528 (at max speed?) each tick.

MoveTo messsages will clear the queue but I don't think MoveTo is meant for use with players(client owned entities)

Spamming messages too fast will block the buffer and stop any other messages coming through

## ClientUnitBehavior.unsynced_client_side_updates_maybe

Currently working on ClientUnitBehavior.unsynced_client_side_updates_maybe, this value is incremented each
time the client sends a movement message, I need to find out how to decrement it.

Sending a MoveTo action can clear it but this is not ideal as MoveTo uses it's own position information and will 
consider the move start point to always be the last MoveTo (ignoring player controlled movements).
This leads me to believe that MoveTo is not intended for client owned entities.

Potentially I need to send a response for each of the move messages but I do not know what message to send.

//`PathManager::RequestPathSync<ClientUnitBehavior>(VectorType const &,VectorType const &,BuildPathSettings const &,ClientUnitBehavior *,void (ClientUnitBehavior::*)(int,Pathfinder *))	.text	00519920	000001F7	0000003C	00000018	R	.	.	.	.	.	.	.`
//PathManager may be related

Sending the position is able to decrement the unsynced positions, but the two unk values must be set correctly or they will be ignored.

`ClientUnitBehavior::OnUpdateApplied` will decrement the counter.

Some of the `i` values below worked but not all

```
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(ClientEntityChannel))
	body.WriteByte(0x35)   // ComponentUpdate
	body.WriteUInt16(0x05) // ComponentID
	body.WriteByte(0x65)   // UnitMoverUpdate

	// UnitBehavior::processUpdate
	body.WriteByte(0xFF) // Unk
	body.WriteByte(0x01) // Update count

	// UnitMoverUpdate::Read
	body.WriteByte(i) // Unk

	body.WriteInt32(p.Position.X)
	body.WriteInt32(p.Position.Y)
	body.WriteInt32(p.Position.Z)

	body.WriteByte(0x02)
	body.WriteUInt32(uint32(p.ClientUpdateNumber))

	//AddSynch(p.Conn, body)

	AddEntityUpdateStreamEnd(body)

	p.send(body)
	i++
```


## Tests


level 16 test:

// Top
00 00 01 00 || 65536
00 00 01 00 || 65536

// Bottom
00 00 01 00 || 65536
00 08 00 00 || 2048

Potentially grid positions are `posX << 10` `posY << 10`, 1 = 0x400


Rotation:
highest value seen 0x16700


Movement example in single message + Header messages

```
Level16 movement from x: 00 00 01 00 y: 00 84 00 00 in the x+ direction mouse

unhandled client entity sub message 3
=======================================================================
unhandled channel message chan: 7 type: 34

time=2021-06-05T19:47:29+01:00 level=info msg=Uncompressed E:
00000000  07 34 05 00 03 25                                 |.4...%|

=======================================================================
<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 22
Player move 0x0 flags? 0x12c00 (%!f(int32=65535), %!f(int32=33791)) Hex (ffff, 83ff)
00000000  07 34 05 00 65 22 01 00  00 2c 01 00 ff ff 00 00  |.4..e"...,......|
00000010  ff 83 00 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 22
Player move 0x0 flags? 0x11400 (%!f(int32=66063), %!f(int32=33844)) Hex (1020f, 8434)
00000000  07 34 05 00 65 22 01 00  00 14 01 00 0f 02 01 00  |.4..e"..........|
00000010  34 84 00 00                                       |4...|

<---- recv [ClientEntityChannel-0x34] len 345
Received 26 player moves unk val: 22
Player move 0x0 flags? 0x11000 (%!f(int32=66591), %!f(int32=33858)) Hex (1041f, 8442)
Player move 0x2 flags? 0x11000 (%!f(int32=67119), %!f(int32=33872)) Hex (1062f, 8450)
Player move 0x0 flags? 0x11000 (%!f(int32=67647), %!f(int32=33886)) Hex (1083f, 845e)
Player move 0x0 flags? 0x11000 (%!f(int32=68175), %!f(int32=33900)) Hex (10a4f, 846c)
Player move 0x0 flags? 0x11000 (%!f(int32=68703), %!f(int32=33914)) Hex (10c5f, 847a)
Player move 0x0 flags? 0x11000 (%!f(int32=69231), %!f(int32=33928)) Hex (10e6f, 8488)
Player move 0x0 flags? 0x11000 (%!f(int32=69759), %!f(int32=33942)) Hex (1107f, 8496)
Player move 0x0 flags? 0x11000 (%!f(int32=70287), %!f(int32=33956)) Hex (1128f, 84a4)
Player move 0x0 flags? 0x11000 (%!f(int32=70815), %!f(int32=33970)) Hex (1149f, 84b2)
Player move 0x0 flags? 0x11000 (%!f(int32=71343), %!f(int32=33984)) Hex (116af, 84c0)
Player move 0x0 flags? 0x11100 (%!f(int32=71871), %!f(int32=34010)) Hex (118bf, 84da)
Player move 0x2 flags? 0x11100 (%!f(int32=72399), %!f(int32=34036)) Hex (11acf, 84f4)
Player move 0x0 flags? 0x11100 (%!f(int32=72927), %!f(int32=34062)) Hex (11cdf, 850e)
Player move 0x0 flags? 0x11100 (%!f(int32=73455), %!f(int32=34088)) Hex (11eef, 8528)
Player move 0x0 flags? 0x11100 (%!f(int32=73983), %!f(int32=34114)) Hex (120ff, 8542)
Player move 0x0 flags? 0x11100 (%!f(int32=74511), %!f(int32=34140)) Hex (1230f, 855c)
Player move 0x0 flags? 0x11100 (%!f(int32=75039), %!f(int32=34166)) Hex (1251f, 8576)
Player move 0x0 flags? 0x11100 (%!f(int32=75567), %!f(int32=34192)) Hex (1272f, 8590)
Player move 0x0 flags? 0x11100 (%!f(int32=76095), %!f(int32=34218)) Hex (1293f, 85aa)
Player move 0x0 flags? 0x11100 (%!f(int32=76623), %!f(int32=34244)) Hex (12b4f, 85c4)
Player move 0x0 flags? 0x11100 (%!f(int32=77151), %!f(int32=34270)) Hex (12d5f, 85de)
Player move 0x0 flags? 0x11200 (%!f(int32=77679), %!f(int32=34303)) Hex (12f6f, 85ff)
Player move 0x2 flags? 0x11200 (%!f(int32=78207), %!f(int32=34336)) Hex (1317f, 8620)
Player move 0x0 flags? 0x11200 (%!f(int32=78735), %!f(int32=34369)) Hex (1338f, 8641)
Player move 0x0 flags? 0x11200 (%!f(int32=79086), %!f(int32=34411)) Hex (134ee, 866b)
Player move 0x1 flags? 0x11200 (%!f(int32=79086), %!f(int32=34411)) Hex (134ee, 866b)
00000000  07 34 05 00 65 22 1a 00  00 10 01 00 1f 04 01 00  |.4..e"..........|
00000010  42 84 00 00 02 00 10 01  00 2f 06 01 00 50 84 00  |B......../...P..|
00000020  00 00 00 10 01 00 3f 08  01 00 5e 84 00 00 00 00  |......?...^.....|
00000030  10 01 00 4f 0a 01 00 6c  84 00 00 00 00 10 01 00  |...O...l........|
00000040  5f 0c 01 00 7a 84 00 00  00 00 10 01 00 6f 0e 01  |_...z........o..|
00000050  00 88 84 00 00 00 00 10  01 00 7f 10 01 00 96 84  |................|
00000060  00 00 00 00 10 01 00 8f  12 01 00 a4 84 00 00 00  |................|
00000070  00 10 01 00 9f 14 01 00  b2 84 00 00 00 00 10 01  |................|
00000080  00 af 16 01 00 c0 84 00  00 00 00 11 01 00 bf 18  |................|
00000090  01 00 da 84 00 00 02 00  11 01 00 cf 1a 01 00 f4  |................|
000000a0  84 00 00 00 00 11 01 00  df 1c 01 00 0e 85 00 00  |................|
000000b0  00 00 11 01 00 ef 1e 01  00 28 85 00 00 00 00 11  |.........(......|
000000c0  01 00 ff 20 01 00 42 85  00 00 00 00 11 01 00 0f  |... ..B.........|
000000d0  23 01 00 5c 85 00 00 00  00 11 01 00 1f 25 01 00  |#..\.........%..|
000000e0  76 85 00 00 00 00 11 01  00 2f 27 01 00 90 85 00  |v......../'.....|
000000f0  00 00 00 11 01 00 3f 29  01 00 aa 85 00 00 00 00  |......?)........|
00000100  11 01 00 4f 2b 01 00 c4  85 00 00 00 00 11 01 00  |...O+...........|
00000110  5f 2d 01 00 de 85 00 00  00 00 12 01 00 6f 2f 01  |_-...........o/.|
00000120  00 ff 85 00 00 02 00 12  01 00 7f 31 01 00 20 86  |...........1.. .|
00000130  00 00 00 00 12 01 00 8f  33 01 00 41 86 00 00 00  |........3..A....|
00000140  00 12 01 00 ee 34 01 00  6b 86 00 00 01 00 12 01  |.....4..k.......|
00000150  00 ee 34 01 00 6b 86 00  00                       |..4..k...|
```

```
Level16 movement from x: 00 00 01 00 y: 00 00 01 00 in the y+ direction mouse

unhandled client entity sub message 3
=======================================================================
unhandled channel message chan: 7 type: 34

time=2021-06-05T19:49:58+01:00 level=info msg=Uncompressed E:
00000000  07 34 05 00 03 2d                                 |.4...-|

=======================================================================
<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 2a
Player move 0x0 flags? 0x200 (%!f(int32=65537), %!f(int32=65536)) Hex (10001, 10000)
00000000  07 34 05 00 65 2a 01 00  00 02 00 00 01 00 01 00  |.4..e*..........|
00000010  00 00 01 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 2a
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=66068)) Hex (10001, 10214)
00000000  07 34 05 00 65 2a 01 00  00 00 00 00 01 00 01 00  |.4..e*..........|
00000010  14 02 01 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 345
Received 26 player moves unk val: 2a
Player move 0x2 flags? 0x0 (%!f(int32=65537), %!f(int32=66600)) Hex (10001, 10428)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=67132)) Hex (10001, 1063c)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=67664)) Hex (10001, 10850)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=68196)) Hex (10001, 10a64)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=68728)) Hex (10001, 10c78)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=69260)) Hex (10001, 10e8c)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=69792)) Hex (10001, 110a0)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=70324)) Hex (10001, 112b4)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=70856)) Hex (10001, 114c8)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=71388)) Hex (10001, 116dc)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=71920)) Hex (10001, 118f0)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=72452)) Hex (10001, 11b04)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=72984)) Hex (10001, 11d18)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=73516)) Hex (10001, 11f2c)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=74048)) Hex (10001, 12140)
Player move 0x0 flags? 0x0 (%!f(int32=65537), %!f(int32=74580)) Hex (10001, 12354)
Player move 0x0 flags? 0x100 (%!f(int32=65528), %!f(int32=75108)) Hex (fff8, 12564)
Player move 0x2 flags? 0x100 (%!f(int32=65519), %!f(int32=75636)) Hex (ffef, 12774)
Player move 0x0 flags? 0x100 (%!f(int32=65510), %!f(int32=76164)) Hex (ffe6, 12984)
Player move 0x0 flags? 0x100 (%!f(int32=65501), %!f(int32=76692)) Hex (ffdd, 12b94)
Player move 0x0 flags? 0x100 (%!f(int32=65492), %!f(int32=77220)) Hex (ffd4, 12da4)
Player move 0x0 flags? 0x100 (%!f(int32=65483), %!f(int32=77748)) Hex (ffcb, 12fb4)
Player move 0x0 flags? 0x100 (%!f(int32=65474), %!f(int32=78276)) Hex (ffc2, 131c4)
Player move 0x0 flags? 0x100 (%!f(int32=65465), %!f(int32=78804)) Hex (ffb9, 133d4)
Player move 0x0 flags? 0x100 (%!f(int32=65442), %!f(int32=79326)) Hex (ffa2, 135de)
Player move 0x1 flags? 0x100 (%!f(int32=65442), %!f(int32=79326)) Hex (ffa2, 135de)
00000000  07 34 05 00 65 2a 1a 02  00 00 00 00 01 00 01 00  |.4..e*..........|
00000010  28 04 01 00 00 00 00 00  00 01 00 01 00 3c 06 01  |(............<..|
00000020  00 00 00 00 00 00 01 00  01 00 50 08 01 00 00 00  |..........P.....|
00000030  00 00 00 01 00 01 00 64  0a 01 00 00 00 00 00 00  |.......d........|
00000040  01 00 01 00 78 0c 01 00  00 00 00 00 00 01 00 01  |....x...........|
00000050  00 8c 0e 01 00 00 00 00  00 00 01 00 01 00 a0 10  |................|
00000060  01 00 00 00 00 00 00 01  00 01 00 b4 12 01 00 00  |................|
00000070  00 00 00 00 01 00 01 00  c8 14 01 00 00 00 00 00  |................|
00000080  00 01 00 01 00 dc 16 01  00 00 00 00 00 00 01 00  |................|
00000090  01 00 f0 18 01 00 00 00  00 00 00 01 00 01 00 04  |................|
000000a0  1b 01 00 00 00 00 00 00  01 00 01 00 18 1d 01 00  |................|
000000b0  00 00 00 00 00 01 00 01  00 2c 1f 01 00 00 00 00  |.........,......|
000000c0  00 00 01 00 01 00 40 21  01 00 00 00 00 00 00 01  |......@!........|
000000d0  00 01 00 54 23 01 00 00  00 01 00 00 f8 ff 00 00  |...T#...........|
000000e0  64 25 01 00 02 00 01 00  00 ef ff 00 00 74 27 01  |d%...........t'.|
000000f0  00 00 00 01 00 00 e6 ff  00 00 84 29 01 00 00 00  |...........)....|
00000100  01 00 00 dd ff 00 00 94  2b 01 00 00 00 01 00 00  |........+.......|
00000110  d4 ff 00 00 a4 2d 01 00  00 00 01 00 00 cb ff 00  |.....-..........|
00000120  00 b4 2f 01 00 00 00 01  00 00 c2 ff 00 00 c4 31  |../............1|
00000130  01 00 00 00 01 00 00 b9  ff 00 00 d4 33 01 00 00  |............3...|
00000140  00 01 00 00 a2 ff 00 00  de 35 01 00 01 00 01 00  |.........5......|
00000150  00 a2 ff 00 00 de 35 01  00                       |......5..|
```

```
Level16 short movement from x: 00 00 01 00 y: 00 00 01 00 in the y+ direction mouse

unhandled client entity sub message 3
=======================================================================
unhandled channel message chan: 7 type: 34

time=2021-06-05T19:53:54+01:00 level=info msg=Uncompressed E:
00000000  07 34 05 00 03 06                                 |.4....|

=======================================================================
<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 3
Player move 0x0 flags? 0x12800 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
00000000  07 34 05 00 65 03 01 00  00 28 01 00 ff ff 00 00  |.4..e....(......|
00000010  ff ff 00 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 3
Player move 0x0 flags? 0x12800 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
00000000  07 34 05 00 65 03 01 00  00 28 01 00 ff ff 00 00  |.4..e....(......|
00000010  ff ff 00 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 3
Player move 0x0 flags? 0x14000 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
00000000  07 34 05 00 65 03 01 00  00 40 01 00 ff ff 00 00  |.4..e....@......|
00000010  ff ff 00 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 137
Received 10 player moves unk val: 3
Player move 0x0 flags? 0x15800 (%!f(int32=65678), %!f(int32=66046)) Hex (1008e, 101fe)
Player move 0x0 flags? 0x700 (%!f(int32=65613), %!f(int32=66573)) Hex (1004d, 1040d)
Player move 0x2 flags? 0x700 (%!f(int32=65548), %!f(int32=67100)) Hex (1000c, 1061c)
Player move 0x0 flags? 0x800 (%!f(int32=65475), %!f(int32=67625)) Hex (ffc3, 10829)
Player move 0x2 flags? 0x800 (%!f(int32=65402), %!f(int32=68150)) Hex (ff7a, 10a36)
Player move 0x0 flags? 0x800 (%!f(int32=65329), %!f(int32=68675)) Hex (ff31, 10c43)
Player move 0x0 flags? 0x800 (%!f(int32=65256), %!f(int32=69200)) Hex (fee8, 10e50)
Player move 0x0 flags? 0x800 (%!f(int32=65183), %!f(int32=69725)) Hex (fe9f, 1105d)
Player move 0x0 flags? 0x800 (%!f(int32=65132), %!f(int32=70001)) Hex (fe6c, 11171)
Player move 0x1 flags? 0x800 (%!f(int32=65132), %!f(int32=70001)) Hex (fe6c, 11171)
00000000  07 34 05 00 65 03 0a 00  00 58 01 00 8e 00 01 00  |.4..e....X......|
00000010  fe 01 01 00 00 00 07 00  00 4d 00 01 00 0d 04 01  |.........M......|
00000020  00 02 00 07 00 00 0c 00  01 00 1c 06 01 00 00 00  |................|
00000030  08 00 00 c3 ff 00 00 29  08 01 00 02 00 08 00 00  |.......)........|
00000040  7a ff 00 00 36 0a 01 00  00 00 08 00 00 31 ff 00  |z...6........1..|
00000050  00 43 0c 01 00 00 00 08  00 00 e8 fe 00 00 50 0e  |.C............P.|
00000060  01 00 00 00 08 00 00 9f  fe 00 00 5d 10 01 00 00  |...........]....|
00000070  00 08 00 00 6c fe 00 00  71 11 01 00 01 00 08 00  |....l...q.......|
00000080  00 6c fe 00 00 71 11 01  00                       |.l...q...|
```

```
Level16 long movement from x: 00 00 01 00 y: 00 00 01 00 in the y+ direction mouse

This movement never finished

unhandled client entity sub message 3
=======================================================================
unhandled channel message chan: 7 type: 34

time=2021-06-05T20:07:48+01:00 level=info msg=Uncompressed E:
00000000  07 34 05 00 03 1f                                 |.4....|

=======================================================================
<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 1c
Player move 0x0 flags? 0x12800 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
00000000  07 34 05 00 65 1c 01 00  00 28 01 00 ff ff 00 00  |.4..e....(......|
00000010  ff ff 00 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: 1c
Player move 0x0 flags? 0x14000 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
00000000  07 34 05 00 65 1c 01 00  00 40 01 00 ff ff 00 00  |.4..e....@......|
00000010  ff ff 00 00                                       |....|

<---- recv [ClientEntityChannel-0x34] len 579
Received 44 player moves unk val: 1c
Player move 0x0 flags? 0x15800 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
Player move 0x0 flags? 0x800 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
Player move 0x0 flags? 0x2000 (%!f(int32=65535), %!f(int32=65535)) Hex (ffff, ffff)
Player move 0x0 flags? 0x3800 (%!f(int32=65094), %!f(int32=65830)) Hex (fe46, 10126)
Player move 0x0 flags? 0x4800 (%!f(int32=64589), %!f(int32=65992)) Hex (fc4d, 101c8)
Player move 0x2 flags? 0x4800 (%!f(int32=64084), %!f(int32=66154)) Hex (fa54, 1026a)
Player move 0x0 flags? 0x4800 (%!f(int32=63579), %!f(int32=66316)) Hex (f85b, 1030c)
Player move 0x0 flags? 0x4800 (%!f(int32=63074), %!f(int32=66478)) Hex (f662, 103ae)
Player move 0x0 flags? 0x4800 (%!f(int32=62569), %!f(int32=66640)) Hex (f469, 10450)
Player move 0x0 flags? 0x4800 (%!f(int32=62064), %!f(int32=66802)) Hex (f270, 104f2)
Player move 0x0 flags? 0x4800 (%!f(int32=61559), %!f(int32=66964)) Hex (f077, 10594)
Player move 0x0 flags? 0x4800 (%!f(int32=61054), %!f(int32=67126)) Hex (ee7e, 10636)
Player move 0x0 flags? 0x4800 (%!f(int32=60549), %!f(int32=67288)) Hex (ec85, 106d8)
Player move 0x0 flags? 0x4800 (%!f(int32=60044), %!f(int32=67450)) Hex (ea8c, 1077a)
Player move 0x0 flags? 0x4800 (%!f(int32=59539), %!f(int32=67612)) Hex (e893, 1081c)
Player move 0x0 flags? 0x4800 (%!f(int32=59034), %!f(int32=67774)) Hex (e69a, 108be)
Player move 0x0 flags? 0x4800 (%!f(int32=58529), %!f(int32=67936)) Hex (e4a1, 10960)
Player move 0x0 flags? 0x4800 (%!f(int32=58024), %!f(int32=68098)) Hex (e2a8, 10a02)
Player move 0x0 flags? 0x4800 (%!f(int32=57519), %!f(int32=68260)) Hex (e0af, 10aa4)
Player move 0x0 flags? 0x4800 (%!f(int32=57014), %!f(int32=68422)) Hex (deb6, 10b46)
Player move 0x0 flags? 0x4800 (%!f(int32=56509), %!f(int32=68584)) Hex (dcbd, 10be8)
Player move 0x0 flags? 0x4800 (%!f(int32=56004), %!f(int32=68746)) Hex (dac4, 10c8a)
Player move 0x0 flags? 0x4800 (%!f(int32=55499), %!f(int32=68908)) Hex (d8cb, 10d2c)
Player move 0x0 flags? 0x4800 (%!f(int32=54994), %!f(int32=69070)) Hex (d6d2, 10dce)
Player move 0x0 flags? 0x4800 (%!f(int32=54489), %!f(int32=69232)) Hex (d4d9, 10e70)
Player move 0x0 flags? 0x4800 (%!f(int32=53984), %!f(int32=69394)) Hex (d2e0, 10f12)
Player move 0x0 flags? 0x4800 (%!f(int32=53479), %!f(int32=69556)) Hex (d0e7, 10fb4)
Player move 0x0 flags? 0x4800 (%!f(int32=52974), %!f(int32=69718)) Hex (ceee, 11056)
Player move 0x0 flags? 0x4800 (%!f(int32=52469), %!f(int32=69880)) Hex (ccf5, 110f8)
Player move 0x0 flags? 0x4800 (%!f(int32=51964), %!f(int32=70042)) Hex (cafc, 1119a)
Player move 0x0 flags? 0x4800 (%!f(int32=51459), %!f(int32=70204)) Hex (c903, 1123c)
Player move 0x0 flags? 0x4800 (%!f(int32=50954), %!f(int32=70366)) Hex (c70a, 112de)
Player move 0x0 flags? 0x4800 (%!f(int32=50449), %!f(int32=70528)) Hex (c511, 11380)
Player move 0x0 flags? 0x4800 (%!f(int32=49944), %!f(int32=70690)) Hex (c318, 11422)
Player move 0x0 flags? 0x4800 (%!f(int32=49439), %!f(int32=70852)) Hex (c11f, 114c4)
Player move 0x0 flags? 0x4800 (%!f(int32=48934), %!f(int32=71014)) Hex (bf26, 11566)
Player move 0x0 flags? 0x4800 (%!f(int32=48429), %!f(int32=71176)) Hex (bd2d, 11608)
Player move 0x0 flags? 0x4800 (%!f(int32=47924), %!f(int32=71338)) Hex (bb34, 116aa)
Player move 0x0 flags? 0x4800 (%!f(int32=47419), %!f(int32=71500)) Hex (b93b, 1174c)
Player move 0x0 flags? 0x4800 (%!f(int32=46914), %!f(int32=71662)) Hex (b742, 117ee)
Player move 0x0 flags? 0x4800 (%!f(int32=46409), %!f(int32=71824)) Hex (b549, 11890)
Player move 0x0 flags? 0x4800 (%!f(int32=45904), %!f(int32=71986)) Hex (b350, 11932)
Player move 0x0 flags? 0x4800 (%!f(int32=45399), %!f(int32=72148)) Hex (b157, 119d4)
Player move 0x0 flags? 0x4800 (%!f(int32=44894), %!f(int32=72310)) Hex (af5e, 11a76)
00000000  07 34 05 00 65 1c 2c 00  00 58 01 00 ff ff 00 00  |.4..e.,..X......|
00000010  ff ff 00 00 00 00 08 00  00 ff ff 00 00 ff ff 00  |................|
00000020  00 00 00 20 00 00 ff ff  00 00 ff ff 00 00 00 00  |... ............|
00000030  38 00 00 46 fe 00 00 26  01 01 00 00 00 48 00 00  |8..F...&.....H..|
00000040  4d fc 00 00 c8 01 01 00  02 00 48 00 00 54 fa 00  |M.........H..T..|
00000050  00 6a 02 01 00 00 00 48  00 00 5b f8 00 00 0c 03  |.j.....H..[.....|
00000060  01 00 00 00 48 00 00 62  f6 00 00 ae 03 01 00 00  |....H..b........|
00000070  00 48 00 00 69 f4 00 00  50 04 01 00 00 00 48 00  |.H..i...P.....H.|
00000080  00 70 f2 00 00 f2 04 01  00 00 00 48 00 00 77 f0  |.p.........H..w.|
00000090  00 00 94 05 01 00 00 00  48 00 00 7e ee 00 00 36  |........H..~...6|
000000a0  06 01 00 00 00 48 00 00  85 ec 00 00 d8 06 01 00  |.....H..........|
000000b0  00 00 48 00 00 8c ea 00  00 7a 07 01 00 00 00 48  |..H......z.....H|
000000c0  00 00 93 e8 00 00 1c 08  01 00 00 00 48 00 00 9a  |............H...|
000000d0  e6 00 00 be 08 01 00 00  00 48 00 00 a1 e4 00 00  |.........H......|
000000e0  60 09 01 00 00 00 48 00  00 a8 e2 00 00 02 0a 01  |`.....H.........|
000000f0  00 00 00 48 00 00 af e0  00 00 a4 0a 01 00 00 00  |...H............|
00000100  48 00 00 b6 de 00 00 46  0b 01 00 00 00 48 00 00  |H......F.....H..|
00000110  bd dc 00 00 e8 0b 01 00  00 00 48 00 00 c4 da 00  |..........H.....|
00000120  00 8a 0c 01 00 00 00 48  00 00 cb d8 00 00 2c 0d  |.......H......,.|
00000130  01 00 00 00 48 00 00 d2  d6 00 00 ce 0d 01 00 00  |....H...........|
00000140  00 48 00 00 d9 d4 00 00  70 0e 01 00 00 00 48 00  |.H......p.....H.|
00000150  00 e0 d2 00 00 12 0f 01  00 00 00 48 00 00 e7 d0  |...........H....|
00000160  00 00 b4 0f 01 00 00 00  48 00 00 ee ce 00 00 56  |........H......V|
00000170  10 01 00 00 00 48 00 00  f5 cc 00 00 f8 10 01 00  |.....H..........|
00000180  00 00 48 00 00 fc ca 00  00 9a 11 01 00 00 00 48  |..H............H|
00000190  00 00 03 c9 00 00 3c 12  01 00 00 00 48 00 00 0a  |......<.....H...|
000001a0  c7 00 00 de 12 01 00 00  00 48 00 00 11 c5 00 00  |.........H......|
000001b0  80 13 01 00 00 00 48 00  00 18 c3 00 00 22 14 01  |......H......"..|
000001c0  00 00 00 48 00 00 1f c1  00 00 c4 14 01 00 00 00  |...H............|
000001d0  48 00 00 26 bf 00 00 66  15 01 00 00 00 48 00 00  |H..&...f.....H..|
000001e0  2d bd 00 00 08 16 01 00  00 00 48 00 00 34 bb 00  |-.........H..4..|
000001f0  00 aa 16 01 00 00 00 48  00 00 3b b9 00 00 4c 17  |.......H..;...L.|
00000200  01 00 00 00 48 00 00 42  b7 00 00 ee 17 01 00 00  |....H..B........|
00000210  00 48 00 00 49 b5 00 00  90 18 01 00 00 00 48 00  |.H..I.........H.|
00000220  00 50 b3 00 00 32 19 01  00 00 00 48 00 00 57 b1  |.P...2.....H..W.|
00000230  00 00 d4 19 01 00 00 00  48 00 00 5e af 00 00 76  |........H..^...v|
00000240  1a 01 00                                          |...|

```