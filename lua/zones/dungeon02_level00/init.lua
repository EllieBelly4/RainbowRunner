npcs = {
    -- This is not spawnable
    --{
    --    name = "base",
    --    position = Vector3.new(400, -152, 49),
    --    rotation = 249
    --},

   -- {
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
      --  name = "npc.Vendor1",
        --position = Vector3.new(195433, 125813 , 10239 ),
	--	position = Vector3.new(140, 188 , 10 ),
      --  rotation = 190,
      --  behaviour = "world.dungeon02.npc.Vendor1.Merchant"
		--behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
  --  },

    
	
	--{
        --X: 423.285156 Y: -77.488281 Z: 49.914062 Rot: 190.00
      --  name = "QuestGiver_L02",--Billy Ray 
        --position = Vector3.new(195433, 125813 , 10239 ),
		--position = Vector3.new(750, 490 , 40 ),
       -- rotation = 360,
      --  position = Vector3.new(180, 498 , 40 ),
      --  rotation = 277.59 ,
		
		--behaviour = "world.tutorial.npc.HermitVendor.Merchant.Weapons"
		--behaviour = "creatures.summon.blinggnome.base.BlingGnome_NPC.Behavior"
   -- },
	
	
	
	
	
}


obelisk = WorldEntity.new("world.checkpoints.Dungeon02CheckpointEntity")
currentZone:spawn(obelisk, Vector3.new(123, 188, 10), 350)


