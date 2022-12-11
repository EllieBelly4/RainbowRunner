function changeZone(player, zoneName)
    player:changeZone(zoneName)
    print("moving to zone " .. zoneName)
end

--function warp(player, x, y, z)
--    player:getChildByGCNativeType("UnitBehavior"):teleport(
--            Vector3.new(x, y, z)
--    )
--    print("warp time")
--end