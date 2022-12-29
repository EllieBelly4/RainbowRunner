npcs = currentZone:baseConfig():getAllNPCs()

for i, v in ipairs(npcs) do
    print(v:fullGCType())

    local npc = currentZone:loadNPCFromConfigFullGCType(v:fullGCType())

    if npc:getChildByGCNativeType("UnitBehavior") == null then
        behaviour = MonsterBehavior2.new("npc.Base.Behavior")
        npc:addChild(behaviour)
    end

    npc:worldEntityFlags(0x6)

    print(v:position():x() .. " " .. v:position():y() .. " " .. v:position():z())

    currentZone:spawn(npc, v:position(), v:heading())
end

obelisk = CheckpointEntity.new("world.checkpoints.TownCheckpointEntity")
currentZone:spawn(obelisk, Vector3.new(441, -168, 49), 92)

function spawnVendorWeapon2()
    vendor = currentZone:loadNPCFromConfig("vendorweapon2")

    behaviour = MonsterBehavior2.new("npc.Base.Behavior")
    vendor:addChild(behaviour)

    merchant = Merchant.new("world.town.npc.VendorWeapon2.Merchant")

    inventory1 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Weapon", 1)
    inventory2 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Armor", 2)
    inventory3 = MerchantInventory.new("world.town.npc.VendorWeapon2.Merchant.Superior", 3)

    merchant:addChild(inventory1)
    merchant:addChild(inventory2)
    merchant:addChild(inventory3)

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

    currentZone:spawn(vendor, Vector3.new(563, 423, 143), 257)
end

spawnVendorWeapon2()