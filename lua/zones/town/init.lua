npcs = currentZone:baseConfig():getAllNPCs()

for i, v in ipairs(npcs) do
    local npc = currentZone:loadNPCFromConfigFullGCType(v:fullGCType())
    currentZone:spawn(npc, v:position(), v:heading())
end

obelisk = CheckpointEntity.new("world.checkpoints.TownCheckpointEntity")
currentZone:spawn(obelisk, Vector3.new(441, -168, 49), 92)
