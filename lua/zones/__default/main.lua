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

    unitBehaviour = player:getChildByGCNativeType("UnitBehavior")

    spawn = ActionSpawn.new()

    pos = nil

    if waypoints["start"] ~= nil then
        pos = waypoints["start"]:position()
    elseif waypoints["waypoint"] ~= nil then
        pos = waypoints["waypoint"]:position()
    else
        customStart = require("config.custom_start")

        lcZoneName = currentZone:name():lower()

        if customStart.locations[lcZoneName] ~= nil then
            pos = customStart.locations[lcZoneName]
        end
    end

    if pos ~= nil then
        spawn:pos(pos)
        unitBehaviour:executeAction(spawn)
    end
end

return module