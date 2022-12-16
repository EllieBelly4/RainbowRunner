function moveNPC(player)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.TownGuard2")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    currentPos = unitBehav:position()

    --X: 356.222656 Y: -182.144531 Z: 49.914062 Rot: 302.00
    print(currentPos:z())

    unitBehav:warpTo(Vector3.new(356, -182, currentPos:z()))
end

function playAnimation(player, unk0, animationIDSelection, animationID, animationFrames, unk4)
    print(unk0 .. " " .. animationIDSelection .. " " .. animationID .. " " .. animationFrames .. " " .. unk4)

    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.oldman1")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    playAnimation = ActionPlayAnimation.new()
    playAnimation:unk0(tonumber(unk0))

    --AnimationID is offset depending on weapon desc, default is +100
    -- 0xFF = Animation ID 0
    -- 0x00 = Animation ID 0
    -- 0x02 = Animation ID v8 + 5
    -- 0x03 = Animation ID v8 + 6
    -- 0x07 = Animation ID v8 + 50
    -- 0xXX = Animation ID from_client_2 + v8
    playAnimation:animationIDSelectionType(tonumber(animationIDSelection))

    -- Looks like only the first uint16 is used in play animation?
    -- Animation ID maybe
    playAnimation:animationID(tonumber(animationID))
    playAnimation:animationFrames(tonumber(animationFrames))

    playAnimation:unk4(tonumber(unk4))

    unitBehav:executeAction(playAnimation)
end

function spawnAnimation (player, unk0, unk1, unk2)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.TownGuard2")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    print("spawn animation")

    spawnAnimation = ActionSpawnAnimation.new()
    spawnAnimation:dataUnk0(tonumber(unk0))
    spawnAnimation:dataUnk1(tonumber(unk1))

    -- Another delay?
    -- Seems to stop animation after spawning in
    spawnAnimation:dataUnk2(tonumber(unk2))

    unitBehav:executeAction(spawnAnimation)
end