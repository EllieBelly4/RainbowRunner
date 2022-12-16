function moveNPC(player)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.TownGuard2")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    currentPos = unitBehav:position()

    --X: 356.222656 Y: -182.144531 Z: 49.914062 Rot: 302.00
    print(currentPos:z())

    unitBehav:warpTo(Vector3.new(356, -182, currentPos:z()))
end

function playAnimation(player)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.TownGuard2")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    playAnimation = ActionPlayAnimation.new()
    playAnimation:unk0(0x01)
    playAnimation:unk1(0x02)
    playAnimation:unk2(0x03)
    playAnimation:unk3(0x04)

    unitBehav:executeAction(playAnimation)
end