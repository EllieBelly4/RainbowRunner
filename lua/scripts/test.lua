function moveNPC(player)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.TownGuard2")

    print(npc:name())
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    --X: 356.222656 Y: -182.144531 Z: 49.914062 Rot: 302.00
    unitBehav:moveTo(Vector2.new(356, -182))
end