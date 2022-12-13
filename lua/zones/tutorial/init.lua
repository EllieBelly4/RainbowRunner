npcs = {
    -- This is not spawnable
    --{
    --    name = "base",
    --    position = Vector3.new(400, -152, 49),
    --    rotation = 249
    --},

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "HermitVendor",--Sam
        --position = Vector3.new(195433, 125813 , 10239 ),
        position = Vector3.new(708.5, 697.05 , 40 ),
        rotation = 190,
        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "helpernoobosaur00",
        --position = Vector3.new(195433, 125813 , 10239 ),
        position = Vector3.new(732, 484 , 40 ),
        rotation = 360,
        behaviour = "npc.Base.Behavior"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "Adventurer1",--Billy Ray
        --position = Vector3.new(195433, 125813 , 10239 ),
        --position = Vector3.new(750, 490 , 40 ),
        -- rotation = 360,
        position = Vector3.new(817, 498 , 40 ),
        rotation = 277.59 ,

        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "Adventurer2",--Kentucky Bill
        --position = Vector3.new(195433, 125813 , 10239 ),
        --position = Vector3.new(780, 490 , 40 ),
        -- rotation = 180,
        position = Vector3.new(800, 490 , 40 ),
        rotation = 280,

        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "Adventurer3",--LizBeth
        --position = Vector3.new(195433, 125813 , 10239 ),
        --position = Vector3.new(800, 490 , 40 ),
        --rotation = 180,
        position = Vector3.new(780, 490 , 40 ),
        rotation = 280,

        -- position = Vector3.new(840, 540 , 40 ),
        -- rotation = 180,

        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "Guard",--Sir Prize
        --position = Vector3.new(195433, 125813 , 10239 ),
        --position = Vector3.new(820, 490 , 40 ),
        -- rotation = 180,
        position = Vector3.new(837, 699 , 40 ),
        rotation = 180,

        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "Recruiter1",--Sir Lee
        --position = Vector3.new(195433, 125813 , 10239 ),
        --position = Vector3.new(840, 490 , 40 ),
        -- rotation = 180,
        position = Vector3.new(839, 448 , 40 ),
        rotation = 352.63 ,

        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },

    {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
        name = "Recruiter2",--Dedre
        --position = Vector3.new(195433, 125813 , 10239 ),
        position = Vector3.new(840, 540 , 40 ),
        rotation = 180,


        --behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
        --behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
    },
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

obelisk = WorldEntity.new("world.checkpoints.TutorialCheckpointEntity")
currentZone:spawn(obelisk, Vector3.new(757, 289, 40), 346)
--portal = ZonePortal.new("unk1", "unk2")
portal = ZonePortal.new("unk1", "unk2")
portal:target("dungeon00_level01")
portal:width(75)
portal:height(75)
portal:unk4(0xFFFFFFFF)

currentZone:spawn(portal, Vector3.new(959, 721, 55), 180)--X: 245522 Y: 184488 Z: 10239 Rot: 357.65

--ZonePortal_oneway = WorldEntity.new("misc.ZonePortal_oneway")
--currentZone:spawn(ZonePortal_oneway, Vector3.new(757, 289, 40), 346)

function spawnHermitVendor()
    vendor = currentZone:loadNPCFromConfig("HermitVendor")

    behaviour = MonsterBehavior2.new("npc.Base.Behavior")
    vendor:addChild(behaviour)

    merchant = Merchant.new("world.tutorial.npc.HermitVendor.Merchant.Weapons")

    inventory1 = MerchantInventory.new("world.tutorial.npc.HermitVendor.Merchant.Weapons", 1)
    --inventory2 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Armor", 2)
    -- inventory3 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Superior", 3)

    merchant:addChild(inventory1)
    --merchant:addChild(inventory2)
    -- merchant:addChild(inventory3)

    vendor:addChild(merchant)

    vendor:level(100)
    vendor:unitFlags(0xFF)
    vendor:hp(1000)
    vendor:mp(2000)

    vendor:unk10Case(0xFF)

    vendor:unk20CaseEntityID(0xFFFF)
    vendor:unk40Case0(0xFFFF)
    vendor:unk40Case1(0xFFFF)
    vendor:unk40Case2(0xFFFF)

    vendor:unk40Case3(0xFF)
    vendor:unk80Case(0xFF)

    vendor:worldEntityFlags(0x7)

    --currentZone:spawn(vendor, Vector3.new(860, 490 , 40), 180)
end

spawnHermitVendor()