npcs = {
    -- This is not spawnable
    --{
    --    name = "base",
    --    position = Vector3.new(400, -152, 49),
    --    rotation = 249
    --},

   -- {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
       -- name = "HermitVendor",--Sam
        --position = Vector3.new(195433, 125813 , 10239 ),
      --  position = Vector3.new(708.5, 697.05 , 40 ),
       -- rotation = 190,
        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
   -- },


}

for i, v in ipairs(npcs) do
    npc = currentZone:loadNPCFromConfig(v["name"])

    if v["behaviour"] then
        npc:removeChildrenByGCNativeType("UnitBehavior")
        behaviour = MonsterBehavior2.new(v["behaviour"])
        npc:addChild(behaviour)
    elseif npc:getChildByGCNativeType("UnitBehavior") ~= null then
        behaviour = MonsterBehavior2.new("npc.Base.Behavior")
        npc:addChild(behaviour)
    end

    npc:worldEntityFlags(0x7)

    currentZone:spawn(npc, v["position"], v["rotation"])
end

--obelisk = WorldEntity.new("world.checkpoints.TutorialCheckpointEntity")
--currentZone:spawn(obelisk, Vector3.new(757, 289, 35), 346)
--portal = ZonePortal.new("unk1", "unk2")
portal = ZonePortal.new("unk1", "unk2")
portal:target("dungeon00_level03")
portal:width(75)
portal:height(75)
portal:unk4(0xFFFFFFFF)
currentZone:spawn(portal, Vector3.new(559, 560, 30), 180)--X: 143463 Y: 143221 Z: 2559 Rot: 358.65

portal = ZonePortal.new("unk1", "unk2")
portal:target("tutorial")
portal:width(75)
portal:height(75)
portal:unk4(0xFFFFFFFF)
currentZone:spawn(portal, Vector3.new(640, -563, 30), 180)--X: 163938 Y: -144150 Z: 2559 Rot: 177.45



--ZonePortal_oneway = WorldEntity.new("misc.ZonePortal_oneway")
--currentZone:spawn(ZonePortal_oneway, Vector3.new(757, 289, 40), 346)

function spawnHermitVendor()
    vendor = currentZone:loadNPCFromConfig("HermitVendor")

    behaviour = MonsterBehavior2.new("npc.Base.Behavior")
    vendor:addChild(behaviour)

    merchant = Merchant.new(" world.tutorial.npc.HermitVendor.Merchant.Weapons")

    inventory1 = MerchantInventory.new(" world.tutorial.npc.HermitVendor.Merchant.Weapons", 1)
    --inventory2 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Armor", 2)
    -- inventory3 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Superior", 3)

    merchant:addChild(inventory1)
    --merchant:addChild(inventory2)
    -- merchant:addChild(inventory3)

    vendor:addChild(merchant)

    vendor:level(100)
    vendor:flags(0xFF)
    vendor:hp(1000)
    vendor:mp(2000)

    vendor:unitUnk10Case(0xFF)

    vendor:unitUnk20CaseEntityID(0xFFFF)
    vendor:unitUnk40Case0(0xFFFF)
    vendor:unitUnk40Case1(0xFFFF)
    vendor:unitUnk40Case2(0xFFFF)

    vendor:unitUnk40Case3(0xFF)
    vendor:unitUnk80Case(0xFF)

    vendor:worldEntityFlags(0x7)

    --currentZone:spawn(vendor, Vector3.new(860, 490 , 40), 180)
end

spawnHermitVendor()