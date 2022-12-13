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