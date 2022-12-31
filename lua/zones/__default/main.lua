zoneCommon = require("lib.zone_common")

function module.__init()
    zoneCommon.spawnEntities(currentZone)
end

function module.__tick()
    --print("__default tick")
end

function module.__onPlayerEnter(player)
    zoneConf = currentZone:baseConfig()
end

return module