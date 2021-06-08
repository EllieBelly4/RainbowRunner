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


```
Level16 long switch from nosync to sync +x mouse only

Received 1 player moves unk val: ff
Player move 0x0 rotation 0x16700(0.00deg) (166, 19997) Hex (a6, 4e1d)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 67 01 00 a6 00 00 00  |.4..e....g......|
00000010  1d 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x14f00(0.00deg) (166, 19997) Hex (a6, 4e1d)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x13700(0.00deg) (166, 19997) Hex (a6, 4e1d)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x11f00(0.00deg) (673, 20149) Hex (2a1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 4f 01 00 a6 00 00 00  |.4..e....O......|
00000010  1d 4e 00 00 00 00 37 01  00 a6 00 00 00 1d 4e 00  |.N....7.......N.|
00000020  00 00 00 1f 01 00 a1 02  00 00 b5 4e 00 00        |...........N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (1205, 20149) Hex (4b5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x2 rotation 0x10e00(0.00deg) (1737, 20149) Hex (6c9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
player started moving
Player move 0x0 rotation 0x10e00(0.00deg) (2269, 20149) Hex (8dd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 b5 04 00 00  |.4..e...........|
00000010  b5 4e 00 00 02 00 0e 01  00 c9 06 00 00 b5 4e 00  |.N............N.|
00000020  00 00 00 0e 01 00 dd 08  00 00 b5 4e 00 00        |...........N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (2801, 20149) Hex (af1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (3333, 20149) Hex (d05, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (3865, 20149) Hex (f19, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 f1 0a 00 00  |.4..e...........|
00000010  b5 4e 00 00 00 00 0e 01  00 05 0d 00 00 b5 4e 00  |.N............N.|
00000020  00 00 00 0e 01 00 19 0f  00 00 b5 4e 00 00        |...........N..|

<---- recv [ClientEntityChannel-0x34] len 33
Received 2 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (4397, 20149) Hex (112d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (4929, 20149) Hex (1341, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 02 00  00 0e 01 00 2d 11 00 00  |.4..e.......-...|
00000010  b5 4e 00 00 00 00 0e 01  00 41 13 00 00 b5 4e 00  |.N.......A....N.|
00000020  00                                                |.|

<---- recv [ClientEntityChannel-0x34] len 33
Received 2 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (5461, 20149) Hex (1555, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (5993, 20149) Hex (1769, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 02 00  00 0e 01 00 55 15 00 00  |.4..e.......U...|
00000010  b5 4e 00 00 00 00 0e 01  00 69 17 00 00 b5 4e 00  |.N.......i....N.|
00000020  00                                                |.|

<---- recv [ClientEntityChannel-0x34] len 59
Received 4 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (6525, 20149) Hex (197d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (7057, 20149) Hex (1b91, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (7589, 20149) Hex (1da5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (8121, 20149) Hex (1fb9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 04 00  00 0e 01 00 7d 19 00 00  |.4..e.......}...|
00000010  b5 4e 00 00 00 00 0e 01  00 91 1b 00 00 b5 4e 00  |.N............N.|
00000020  00 00 00 0e 01 00 a5 1d  00 00 b5 4e 00 00 00 00  |...........N....|
00000030  0e 01 00 b9 1f 00 00 b5  4e 00 00                 |........N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (8653, 20149) Hex (21cd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (9185, 20149) Hex (23e1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (9717, 20149) Hex (25f5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 cd 21 00 00  |.4..e........!..|
00000010  b5 4e 00 00 00 00 0e 01  00 e1 23 00 00 b5 4e 00  |.N........#...N.|
00000020  00 00 00 0e 01 00 f5 25  00 00 b5 4e 00 00        |.......%...N..|

<---- recv [ClientEntityChannel-0x34] len 33
Received 2 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (10249, 20149) Hex (2809, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (10781, 20149) Hex (2a1d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 02 00  00 0e 01 00 09 28 00 00  |.4..e........(..|
00000010  b5 4e 00 00 00 00 0e 01  00 1d 2a 00 00 b5 4e 00  |.N........*...N.|
00000020  00                                                |.|

<---- recv [ClientEntityChannel-0x34] len 33
Received 2 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (11313, 20149) Hex (2c31, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (11845, 20149) Hex (2e45, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 02 00  00 0e 01 00 31 2c 00 00  |.4..e.......1,..|
00000010  b5 4e 00 00 00 00 0e 01  00 45 2e 00 00 b5 4e 00  |.N.......E....N.|
00000020  00                                                |.|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (12377, 20149) Hex (3059, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (12909, 20149) Hex (326d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (13441, 20149) Hex (3481, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 59 30 00 00  |.4..e.......Y0..|
00000010  b5 4e 00 00 00 00 0e 01  00 6d 32 00 00 b5 4e 00  |.N.......m2...N.|
00000020  00 00 00 0e 01 00 81 34  00 00 b5 4e 00 00        |.......4...N..|

<---- recv [ClientEntityChannel-0x34] len 33
Received 2 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (13973, 20149) Hex (3695, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (14505, 20149) Hex (38a9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 02 00  00 0e 01 00 95 36 00 00  |.4..e........6..|
00000010  b5 4e 00 00 00 00 0e 01  00 a9 38 00 00 b5 4e 00  |.N........8...N.|
00000020  00                                                |.|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (15037, 20149) Hex (3abd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (15569, 20149) Hex (3cd1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (16101, 20149) Hex (3ee5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 bd 3a 00 00  |.4..e........:..|
00000010  b5 4e 00 00 00 00 0e 01  00 d1 3c 00 00 b5 4e 00  |.N........<...N.|
00000020  00 00 00 0e 01 00 e5 3e  00 00 b5 4e 00 00        |.......>...N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (16633, 20149) Hex (40f9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (17165, 20149) Hex (430d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (17697, 20149) Hex (4521, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 f9 40 00 00  |.4..e........@..|
00000010  b5 4e 00 00 00 00 0e 01  00 0d 43 00 00 b5 4e 00  |.N........C...N.|
00000020  00 00 00 0e 01 00 21 45  00 00 b5 4e 00 00        |......!E...N..|

<---- recv [ClientEntityChannel-0x34] len 59
Received 4 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (18229, 20149) Hex (4735, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (18761, 20149) Hex (4949, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (19293, 20149) Hex (4b5d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (19825, 20149) Hex (4d71, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 04 00  00 0e 01 00 35 47 00 00  |.4..e.......5G..|
00000010  b5 4e 00 00 00 00 0e 01  00 49 49 00 00 b5 4e 00  |.N.......II...N.|
00000020  00 00 00 0e 01 00 5d 4b  00 00 b5 4e 00 00 00 00  |......]K...N....|
00000030  0e 01 00 71 4d 00 00 b5  4e 00 00                 |...qM...N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (20357, 20149) Hex (4f85, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (20889, 20149) Hex (5199, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (21421, 20149) Hex (53ad, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 85 4f 00 00  |.4..e........O..|
00000010  b5 4e 00 00 00 00 0e 01  00 99 51 00 00 b5 4e 00  |.N........Q...N.|
00000020  00 00 00 0e 01 00 ad 53  00 00 b5 4e 00 00        |.......S...N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (21953, 20149) Hex (55c1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (22485, 20149) Hex (57d5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (23017, 20149) Hex (59e9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 c1 55 00 00  |.4..e........U..|
00000010  b5 4e 00 00 00 00 0e 01  00 d5 57 00 00 b5 4e 00  |.N........W...N.|
00000020  00 00 00 0e 01 00 e9 59  00 00 b5 4e 00 00        |.......Y...N..|

<---- recv [ClientEntityChannel-0x34] len 59
Received 4 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (23549, 20149) Hex (5bfd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (24081, 20149) Hex (5e11, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (24613, 20149) Hex (6025, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (25145, 20149) Hex (6239, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 04 00  00 0e 01 00 fd 5b 00 00  |.4..e........[..|
00000010  b5 4e 00 00 00 00 0e 01  00 11 5e 00 00 b5 4e 00  |.N........^...N.|
00000020  00 00 00 0e 01 00 25 60  00 00 b5 4e 00 00 00 00  |......%`...N....|
00000030  0e 01 00 39 62 00 00 b5  4e 00 00                 |...9b...N..|

<---- recv [ClientEntityChannel-0x34] len 46
Received 3 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (25677, 20149) Hex (644d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (26209, 20149) Hex (6661, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (26741, 20149) Hex (6875, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 03 00  00 0e 01 00 4d 64 00 00  |.4..e.......Md..|
00000010  b5 4e 00 00 00 00 0e 01  00 61 66 00 00 b5 4e 00  |.N.......af...N.|
00000020  00 00 00 0e 01 00 75 68  00 00 b5 4e 00 00        |......uh...N..|

<---- recv [ClientEntityChannel-0x34] len 59
Received 4 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (27273, 20149) Hex (6a89, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (27805, 20149) Hex (6c9d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (28337, 20149) Hex (6eb1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (28869, 20149) Hex (70c5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 04 00  00 0e 01 00 89 6a 00 00  |.4..e........j..|
00000010  b5 4e 00 00 00 00 0e 01  00 9d 6c 00 00 b5 4e 00  |.N........l...N.|
00000020  00 00 00 0e 01 00 b1 6e  00 00 b5 4e 00 00 00 00  |.......n...N....|
00000030  0e 01 00 c5 70 00 00 b5  4e 00 00                 |....p...N..|

<---- recv [ClientEntityChannel-0x34] len 33
Received 2 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (29401, 20149) Hex (72d9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
Player move 0x0 rotation 0x10e00(0.00deg) (29933, 20149) Hex (74ed, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 02 00  00 0e 01 00 d9 72 00 00  |.4..e........r..|
00000010  b5 4e 00 00 00 00 0e 01  00 ed 74 00 00 b5 4e 00  |.N........t...N.|
00000020  00                                                |.|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (30465, 20149) Hex (7701, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 01 77 00 00  |.4..e........w..|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (30997, 20149) Hex (7915, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 15 79 00 00  |.4..e........y..|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (31529, 20149) Hex (7b29, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 29 7b 00 00  |.4..e.......){..|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (32061, 20149) Hex (7d3d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 3d 7d 00 00  |.4..e.......=}..|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (32593, 20149) Hex (7f51, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 51 7f 00 00  |.4..e.......Q...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (33125, 20149) Hex (8165, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 65 81 00 00  |.4..e.......e...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (33657, 20149) Hex (8379, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 79 83 00 00  |.4..e.......y...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (34189, 20149) Hex (858d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 8d 85 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (34721, 20149) Hex (87a1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 a1 87 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (35253, 20149) Hex (89b5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 b5 89 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (35785, 20149) Hex (8bc9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 c9 8b 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (36317, 20149) Hex (8ddd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 dd 8d 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (36849, 20149) Hex (8ff1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 f1 8f 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (37381, 20149) Hex (9205, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 05 92 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (37913, 20149) Hex (9419, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 19 94 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (38445, 20149) Hex (962d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 2d 96 00 00  |.4..e.......-...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (38977, 20149) Hex (9841, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 41 98 00 00  |.4..e.......A...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (39509, 20149) Hex (9a55, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 55 9a 00 00  |.4..e.......U...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (40041, 20149) Hex (9c69, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 69 9c 00 00  |.4..e.......i...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (40573, 20149) Hex (9e7d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 7d 9e 00 00  |.4..e.......}...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (41105, 20149) Hex (a091, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 91 a0 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (41637, 20149) Hex (a2a5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 a5 a2 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (42169, 20149) Hex (a4b9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 b9 a4 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (42701, 20149) Hex (a6cd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 cd a6 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (43233, 20149) Hex (a8e1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 e1 a8 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (43765, 20149) Hex (aaf5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 f5 aa 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (44297, 20149) Hex (ad09, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 09 ad 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (44829, 20149) Hex (af1d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 1d af 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (45361, 20149) Hex (b131, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 31 b1 00 00  |.4..e.......1...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (45893, 20149) Hex (b345, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 45 b3 00 00  |.4..e.......E...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (46425, 20149) Hex (b559, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 59 b5 00 00  |.4..e.......Y...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (46957, 20149) Hex (b76d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 6d b7 00 00  |.4..e.......m...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (47489, 20149) Hex (b981, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 81 b9 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (48021, 20149) Hex (bb95, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 95 bb 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (48553, 20149) Hex (bda9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 a9 bd 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (49085, 20149) Hex (bfbd, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 bd bf 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (49617, 20149) Hex (c1d1, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 d1 c1 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (50149, 20149) Hex (c3e5, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 e5 c3 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (50681, 20149) Hex (c5f9, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 f9 c5 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (51213, 20149) Hex (c80d, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 0d c8 00 00  |.4..e...........|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (51745, 20149) Hex (ca21, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 21 ca 00 00  |.4..e.......!...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10e00(0.00deg) (52277, 20149) Hex (cc35, 4eb5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0e 01 00 35 cc 00 00  |.4..e.......5...|
00000010  b5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (52805, 20155) Hex (ce45, 4ebb)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 45 ce 00 00  |.4..e.......E...|
00000010  bb 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x2 rotation 0x10f00(0.00deg) (53333, 20161) Hex (d055, 4ec1)
>>>>> send [ClientEntityChannel-53] len 26
player started moving
00000000  07 34 05 00 65 ff 01 02  00 0f 01 00 55 d0 00 00  |.4..e.......U...|
00000010  c1 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (53861, 20167) Hex (d265, 4ec7)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 65 d2 00 00  |.4..e.......e...|
00000010  c7 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (54389, 20173) Hex (d475, 4ecd)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 75 d4 00 00  |.4..e.......u...|
00000010  cd 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (54917, 20179) Hex (d685, 4ed3)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 85 d6 00 00  |.4..e...........|
00000010  d3 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (55445, 20185) Hex (d895, 4ed9)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 95 d8 00 00  |.4..e...........|
00000010  d9 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (55973, 20191) Hex (daa5, 4edf)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 a5 da 00 00  |.4..e...........|
00000010  df 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (56501, 20197) Hex (dcb5, 4ee5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 b5 dc 00 00  |.4..e...........|
00000010  e5 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (57029, 20203) Hex (dec5, 4eeb)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 c5 de 00 00  |.4..e...........|
00000010  eb 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (57557, 20209) Hex (e0d5, 4ef1)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 d5 e0 00 00  |.4..e...........|
00000010  f1 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (58085, 20215) Hex (e2e5, 4ef7)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 e5 e2 00 00  |.4..e...........|
00000010  f7 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (58613, 20221) Hex (e4f5, 4efd)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 f5 e4 00 00  |.4..e...........|
00000010  fd 4e 00 00                                       |.N..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (59141, 20227) Hex (e705, 4f03)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 05 e7 00 00  |.4..e...........|
00000010  03 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (59669, 20233) Hex (e915, 4f09)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 15 e9 00 00  |.4..e...........|
00000010  09 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (60197, 20239) Hex (eb25, 4f0f)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 25 eb 00 00  |.4..e.......%...|
00000010  0f 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (60725, 20245) Hex (ed35, 4f15)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 35 ed 00 00  |.4..e.......5...|
00000010  15 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (61253, 20251) Hex (ef45, 4f1b)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 45 ef 00 00  |.4..e.......E...|
00000010  1b 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (61781, 20257) Hex (f155, 4f21)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 55 f1 00 00  |.4..e.......U...|
00000010  21 4f 00 00                                       |!O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (62309, 20263) Hex (f365, 4f27)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 65 f3 00 00  |.4..e.......e...|
00000010  27 4f 00 00                                       |'O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (62837, 20269) Hex (f575, 4f2d)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 75 f5 00 00  |.4..e.......u...|
00000010  2d 4f 00 00                                       |-O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (63365, 20275) Hex (f785, 4f33)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 85 f7 00 00  |.4..e...........|
00000010  33 4f 00 00                                       |3O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (63893, 20281) Hex (f995, 4f39)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 95 f9 00 00  |.4..e...........|
00000010  39 4f 00 00                                       |9O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (64421, 20287) Hex (fba5, 4f3f)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 a5 fb 00 00  |.4..e...........|
00000010  3f 4f 00 00                                       |?O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (64949, 20293) Hex (fdb5, 4f45)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 b5 fd 00 00  |.4..e...........|
00000010  45 4f 00 00                                       |EO..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (65477, 20299) Hex (ffc5, 4f4b)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 c5 ff 00 00  |.4..e...........|
00000010  4b 4f 00 00                                       |KO..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (66005, 20305) Hex (101d5, 4f51)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 d5 01 01 00  |.4..e...........|
00000010  51 4f 00 00                                       |QO..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x10f00(0.00deg) (66533, 20311) Hex (103e5, 4f57)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 0f 01 00 e5 03 01 00  |.4..e...........|
00000010  57 4f 00 00                                       |WO..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (67061, 20325) Hex (105f5, 4f65)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 f5 05 01 00  |.4..e...........|
00000010  65 4f 00 00                                       |eO..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x2 rotation 0x11000(0.00deg) (67589, 20339) Hex (10805, 4f73)
>>>>> send [ClientEntityChannel-53] len 26
player started moving
00000000  07 34 05 00 65 ff 01 02  00 10 01 00 05 08 01 00  |.4..e...........|
00000010  73 4f 00 00                                       |sO..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (68117, 20353) Hex (10a15, 4f81)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 15 0a 01 00  |.4..e...........|
00000010  81 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (68645, 20367) Hex (10c25, 4f8f)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 25 0c 01 00  |.4..e.......%...|
00000010  8f 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (69173, 20381) Hex (10e35, 4f9d)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 35 0e 01 00  |.4..e.......5...|
00000010  9d 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (69701, 20395) Hex (11045, 4fab)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 45 10 01 00  |.4..e.......E...|
00000010  ab 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (70229, 20409) Hex (11255, 4fb9)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 55 12 01 00  |.4..e.......U...|
00000010  b9 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (70757, 20423) Hex (11465, 4fc7)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 65 14 01 00  |.4..e.......e...|
00000010  c7 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (71285, 20437) Hex (11675, 4fd5)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 75 16 01 00  |.4..e.......u...|
00000010  d5 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11000(0.00deg) (71813, 20451) Hex (11885, 4fe3)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 10 01 00 85 18 01 00  |.4..e...........|
00000010  e3 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11100(0.00deg) (72341, 20477) Hex (11a95, 4ffd)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 11 01 00 95 1a 01 00  |.4..e...........|
00000010  fd 4f 00 00                                       |.O..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x2 rotation 0x11100(0.00deg) (72869, 20503) Hex (11ca5, 5017)
>>>>> send [ClientEntityChannel-53] len 26
player started moving
00000000  07 34 05 00 65 ff 01 02  00 11 01 00 a5 1c 01 00  |.4..e...........|
00000010  17 50 00 00                                       |.P..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11100(0.00deg) (73397, 20529) Hex (11eb5, 5031)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 11 01 00 b5 1e 01 00  |.4..e...........|
00000010  31 50 00 00                                       |1P..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11100(0.00deg) (73925, 20555) Hex (120c5, 504b)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 11 01 00 c5 20 01 00  |.4..e........ ..|
00000010  4b 50 00 00                                       |KP..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x0 rotation 0x11100(0.00deg) (74238, 20586) Hex (121fe, 506a)
>>>>> send [ClientEntityChannel-53] len 26
00000000  07 34 05 00 65 ff 01 00  00 11 01 00 fe 21 01 00  |.4..e........!..|
00000010  6a 50 00 00                                       |jP..|

<---- recv [ClientEntityChannel-0x34] len 20
Received 1 player moves unk val: ff
Player move 0x1 rotation 0x11100(0.00deg) (74238, 20586) Hex (121fe, 506a)
>>>>> send [ClientEntityChannel-53] len 26
player finished moving
00000000  07 34 05 00 65 ff 01 01  00 11 01 00 fe 21 01 00  |.4..e........!..|
00000010  6a 50 00 00                                       |jP..|
```
