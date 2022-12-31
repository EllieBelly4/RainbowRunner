function test(player)
    unitBehaviour = player:getChildByGCNativeType("UnitBehavior")
end

function createWaypoint(player)
    zone = player:zone()
    wp = zone:loadWaypointFromConfig("start")

    wpConf = wp:baseConfig()

    zone:spawn(wp, wpConf:position(), wpConf:heading())
end

function moveNPC(player, x, y)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.oldman1")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    currentPos = unitBehav:position()

    --X: 356.222656 Y: -182.144531 Z: 49.914062 Rot: 302.00
--     print(currentPos:z())

    targetPos = Vector3.new(356, -182, currentPos:z())

    unitBehav:moveTo(Vector2.new(currentPos:x()+tonumber(x),currentPos:y()+tonumber(y)))
end

function playAnimation(player, animationIDSelection, animationID, unk0, unk4)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.oldman1")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")
    animationIDNumber = tonumber(animationID)

    animations = npc:animations()

    if animations == nil then
        print("Entity has no animations")
        return
    end

    playAnimation = ActionPlayAnimation.new()

    if tonumber(animationIDSelection) > 7 then
        animation = nil

        for _, anim in pairs(animations) do
            if anim:id() == animationIDNumber then
                animation = anim
                break
            end
        end

        if animation == nil then
            print("Entity does not have animation with ID " .. animationID)
            return
        end

        print("Playing animation ID: " .. animationID)
        playAnimation:animationID(animationIDNumber)
    end

    playAnimation:animationIDSelectionType(0x32)

    if animationIDSelection ~= nil and animationIDSelection ~= "" then
        playAnimation:animationIDSelectionType(tonumber(animationIDSelection))

        if tonumber(animationIDSelection) == 2 then
            playAnimation:animationFrames(20)
        else
            playAnimation:animationFrames(5)
        end
    else
        playAnimation:animationFrames(animation:numFrames())
    end

    if unk4 ~= nil and unk4 ~= "" then
        playAnimation:unk4(tonumber(unk4))
    end

    if unk0 ~= nil and unk0 ~= "" then
        playAnimation:unk0(tonumber(unk0))
    end

    unitBehav:executeAction(playAnimation)
end

function listAnimations (player)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.oldman1")
    animations = npc:animations()

    if animations == nil then
        print("Entity has no animations")
        return
    end

    resultString = ""

    for _, anim in pairs(animations) do
        resultString = resultString .. anim:id() .. ", "
    end

    print(resultString)
end

function spawnAnimation (player, unk0, unk1, unk2)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.TownGuard2")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    print("spawn animation")

    spawnAnimation = ActionSpawnAnimation.new()
    spawnAnimation:dataUnk0(tonumber(unk0))
    spawnAnimation:dataUnk1(tonumber(unk1))

    -- Another delay?
    -- Seems to stop animation after spawning in
    spawnAnimation:dataUnk2(tonumber(unk2))

    unitBehav:executeAction(spawnAnimation)
end

function move(player, unk0)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.oldman1")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    print("move")

    move = ActionMoveInDirectionAction.new()
    move:unk0(tonumber(unk0))

    unitBehav:executeAction(move)
end

function attack(player, unk0, id)
    zone = player:zone()
    npc = zone:findEntityByGCTypeName("world.town.npc.oldman1")
    unitBehav = npc:getChildByGCNativeType("UnitBehavior")

    print("attack")

    attack = ActionAttackTarget2.new()
    attack:unk0(tonumber(unk0))
    attack:targetID(tonumber(id))

    unitBehav:executeAction(attack)
end

function usePosition(player, actionID)
    zone = player:zone()
    avatar = zone:findEntityByGCTypeName("avatar.classes.fighterfemale")
    unitBehav = avatar:getChildByGCNativeType("UnitBehavior")

    print("use position")

    usePosition = ActionUsePosition.new()
    usePosition:actionID(tonumber(actionID))
    usePosition:positionX()
    usePosition:positionY()
    usePosition:positionZ()

--     unitBehav:executeAction(usePosition)
end

function skillSlot(player, slot)
    zone = player:zone()
    avatar = zone:findEntityByGCTypeName("avatar.classes.fighterfemale")
    unitBehav = avatar:getChildByGCNativeType("UnitBehavior")

    skillsComp = avatar:getChildByGCNativeType("Skills")
    firstSkill = skillsComp:getActiveSkills()[1]

    print(firstSkill:gctype())

    if firstSkill:slot() ~= 0 then
        skillsComp:unequipSkill(firstSkill)
    end

    skillsComp:equipSkill(firstSkill, tonumber(slot))
end
