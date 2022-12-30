npcs = currentZone:baseConfig():getAllNPCs()

for i, v in ipairs(npcs) do
    local npc = currentZone:loadNPCFromConfigFullGCType(v:fullGCType())
    currentZone:spawn(npc, v:position(), v:heading())
end

obelisk = CheckpointEntity.new("world.checkpoints.TutorialCheckpointEntity")
currentZone:spawn(obelisk, Vector3.new(757, 289, 40), 346)