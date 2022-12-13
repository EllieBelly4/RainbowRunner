
obelisk = WorldEntity.new("world.checkpoints.TutorialCheckpointEntity")
currentZone:spawn(obelisk, Vector3.new(757, 289, 35), 346)--X: 143048 Y: 147450 Z: 2559 Rot: 361.16
--portal = ZonePortal.new("unk1", "unk2")
portal = ZonePortal.new("unk1", "unk2")
portal:target("tutorial")
portal:width(75)
portal:height(75)
portal:unk4(0xFFFFFFFF)
currentZone:spawn(portal, Vector3.new(641, -557, 35), 180)--X: 143116 Y: 246097 Z: 2559 Rot: 2.59

portal = ZonePortal.new("unk1", "unk2")
portal:target("dungeon00_level02")
portal:width(75)
portal:height(75)
portal:unk4(0xFFFFFFFF)

currentZone:spawn(portal, Vector3.new(560, 962, 35), 180)--X: 143336 Y: 246261 Z: 2559 Rot: 0.76





--ZonePortal_oneway = WorldEntity.new("misc.ZonePortal_oneway")
--currentZone:spawn(ZonePortal_oneway, Vector3.new(757, 289, 40), 346)



    --currentZone:spawn(vendor, Vector3.new(860, 490 , 40), 180)



mobs =
{
    -- This is not spawnable
    --{
    --    name = "base",
    --    position = Vector3.new(400, -152, 49),
    --    rotation = 249
    --},

   {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "RattleTooth",--Sam
        --position = Vector3.new(195433, 125813 , 10239 ),
       position = Vector3.new(514, 357 , 40 ),--X: 131687 Y: -91364 Z: 2559 Rot: 214.34
       rotation = 214,
       -- behaviour = "world.dungeon00.mob"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
   },


}

for i, v in ipairs(mobs) do
    mob = currentZone:loadMOBFromConfig(v["name"])

    if v["behaviour"] then
        mob:removeChildrenByGCNativeType("UnitBehavior")
        behaviour = MonsterBehavior2.new(v["behaviour"])
        mob:addChild(behaviour)
    elseif mob:getChildByGCNativeType("UnitBehavior") ~= null then
        behaviour = MonsterBehavior2.new("mob.Base.Behavior")
        mob:addChild(behaviour)
    end

    mob:worldEntityFlags(0x7)

    currentZone:spawn(mob, v["position"], v["rotation"])
end



spawnMob()