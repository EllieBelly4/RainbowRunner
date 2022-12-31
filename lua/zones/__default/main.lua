zoneCommon = require("lib.zone_common")

function module.__init()
    zoneCommon.spawnEntities(currentZone)
end

function module.__tick()
    --print("__default tick")
end

function module.__onPlayerEnter(player)
    zoneConf = currentZone:baseConfig()
    waypoints = zoneConf:waypoints()

    startWaypoint = waypoints["start"]

    if startWaypoint ~= nil then
        unitBehaviour = player:getChildByGCNativeType("UnitBehavior")

        print(startWaypoint:position():string())

        unitBehaviour:warpTo(startWaypoint:position())
    end
end

return module