function changeZone(player, zoneName)
    player:changeZone(zoneName)
    print("moving to zone " .. zoneName)
end

function warp(player, x, y, z)
    avatar = player:getChildByGCNativeType("Avatar")

    if avatar == nil then
        print("no avatar")
        return
    end

    avatar:teleport(
            Vector3.new(tonumber(x), tonumber(y), tonumber(z))
    )

    print("warping to " .. x .. ", " .. y .. ", " .. z)
end

function warpRelative(player, x, y, z)
    avatar = player:getChildByGCNativeType("Avatar")

    if avatar == nil then
        print("no avatar")
        return
    end

    unitBehav = avatar:getChildByGCNativeType("UnitBehavior")

    if unitBehav == nil then
        print("no unit behaviour")
        return
    end

    position = unitBehav:position()

    newPos = Vector3.new(position:x() + tonumber(x), position:y() + tonumber(y), position:z() + tonumber(z))

    avatar:teleport(newPos)

    print("warping to " .. newPos:x() .. ", " .. newPos:y() .. ", " .. newPos:z())
end