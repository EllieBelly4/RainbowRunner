zoneCommon = require("lib.zone_common")

function module.__init()
    zoneCommon.spawnEntities(currentZone)
end

function module.__tick()
end

function module.__onPlayerEnter(player)
    require("zones.__default.main").__onPlayerEnter(player)
end

return module