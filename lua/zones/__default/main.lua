zoneCommon = require("lib.zone_common")

function module.__init()
--     zoneCommon.spawnEntities(currentZone)
end

function module.__tick()
    --print("__default tick")
end

function module.__onPlayerEnter(player)
    zoneConf = currentZone:baseConfig()
    waypoints = zoneConf:waypoints()

    unitBehaviour = player:getChildByGCNativeType("UnitBehavior")

    if waypoints["start"] ~= nil then
        unitBehaviour:warpTo(waypoints["start"]:position())
    elseif waypoints["waypoint"] ~= nil then
        unitBehaviour:warpTo(waypoints["waypoint"]:position())
    else
        customStart = require("config.custom_start")

        lcZoneName = currentZone:name():lower()

        if customStart.locations[lcZoneName] ~= nil then
            unitBehaviour:warpTo(customStart.locations[lcZoneName])
        end
    end
end

return module