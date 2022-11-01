--obelisk = WorldEntity.new("world.checkpoints.TownCheckpointEntity")
--currentZone:spawn(obelisk, Vector3.new(441, -168, 49), 92)

--snowman = currentZone:loadNPCFromConfig("snowman1")
--currentZone:spawn(snowman, Vector3.new(415, -168, 49), 92)

--snowman = currentZone:loadNPCFromConfig("amazon1")
--currentZone:spawn(snowman, Vector3.new(415, -168, 49), 92)

noobosaur = currentZone:loadNPCFromConfig("helpernoobosaur01")

removed = noobosaur:removeChildrenByGCNativeType("UnitBehavior")

if removed == 0 then
    print("did not remove behavior")
end

noobosaurBehaviour = MonsterBehavior2.new("npc.Base.Behavior")
noobosaur:addChild(noobosaurBehaviour)

currentZone:spawn(noobosaur, Vector3.new(415, -168, 49), 92)