function toggleMovementMessages(player)
    rrplayer = player:getRRPlayer()

    old = rrplayer:getDebugSendMovementMessages()

    rrplayer:setDebugSendMovementMessages(not old)

    print("Movement messages are now " .. tostring(not old))
end