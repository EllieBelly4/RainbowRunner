-- local path =

local currentEntity
local unitBehaviour
local path = {
    Vector2.new(441.050781, -211.964844),
    Vector2.new(499.093750, -171.375000),
    Vector2.new(438.953125, -110.675781),
    Vector2.new(387.273438, -163.656250),
}
local startTime

local pathIndex = 1

local pathReachTime = 0
local wasMoving = false

function module.__init(entity)
    currentEntity = entity
    unitBehaviour = entity:getChildByGCNativeType("UnitBehavior")
    startTime = Time.time()
end

function module.__tick()
    if Time.time() - startTime < 1 then
        return
    end

    if unitBehaviour:isMoving() then
        return
    end

    if wasMoving then
        pathReachTime = Time.time()
    end

    wasMoving = unitBehaviour:isMoving()

    -- This is to account for turn speed, which for this character is 360 degrees per second
    if Time.time() < pathReachTime + 1.1 then
        return
    end

--     print("next node ".. pathIndex)

    unitBehaviour:moveTo(path[pathIndex])

    pathIndex = pathIndex + 1
    if pathIndex > #path then
        pathIndex = 1
    end
end

return module