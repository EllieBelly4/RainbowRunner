function module.spawnEntities(zone)
    for i, v in ipairs(zone:baseConfig():getAllNPCs()) do
        local entity = currentZone:loadNPCFromConfig(v:name())
        currentZone:spawn(entity, v:position(), v:heading())
    end

    for i, v in ipairs(zone:baseConfig():getAllCheckpointEntities()) do
        local entity = currentZone:loadCheckpointEntityFromConfig(v:name())
        currentZone:spawn(entity, v:position(), v:heading())
    end
end

return module