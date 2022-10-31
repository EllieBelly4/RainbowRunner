print(currentZone:name())

npc1 = npc.new("world.town.npc.Amazon1")
npc1:name("Vendor_Amazon1")

currentZone:spawnNPC(npc1)

print(npc1:name())