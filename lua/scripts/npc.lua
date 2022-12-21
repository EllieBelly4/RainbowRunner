function spawnZoneNPC(player, name, x, y, z, rotation)
    zone = player:zone()
    npc = zone:loadNPCFromConfig(name)

    if npc:getChildByGCNativeType("UnitBehavior") ~= null then
        behaviour = MonsterBehavior2.new("npc.Base.Behavior")
        npc:addChild(behaviour)
    end

    npc:worldEntityFlags(0x7)

    currentZone:spawn(npc, Vector3.new(x, y, z), rotation)
end

function test()
    print("test")
end