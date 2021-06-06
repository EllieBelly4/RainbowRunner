```
# UnitBehavior - UnitMoverUpdate:read
35 # ComponentUpdate
05 00 # Component ID
65 # Command UnitMoverUpdate::read
05 # Unk UnitBehavior::processUpdate
01 # Unk UnitBehavior::processUpdate Update count
06 # Unk
02 02 02 02 # PosX?
03 03 03 03 # PosY?
04 04 04 04 # PosZ?
02 00 00 00 00 # Synch
06 # End                               
```


```
# UnitBehavior - UnitMoverUpdate::read
35 # ComponentUpdate
05 00 # Component ID
# Command
# 05 - Behavior::terminateAllActionsLocal
# 65 - UnitMoverUpdate::read
65 # Command 
05 # Unk UnitBehavior::processUpdate
01 # Unk UnitBehavior::processUpdate Update count
06 # Unk
10 10 00 00 # PosX?
00 10 10 00 # PosY?
00 10 00 00 # PosZ?
02 00 7e 04 00 # Synch
06 # End  
```

```
# UnitBehavior - Activate::readData
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
06 # ActionID

# Activate::readData
01
05 00

02 00 7e 04 00 # Synch
06 # End   
```

```
# UnitBehavior - MoveTo::readData
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
01 # ActionID

# MoveTo::readData
FF
7f 08 00 00
7f 08 00 00
#00 01 00 00
#00 00 01 00

02 00 00 00 00 # Synch
06 # End
```

```
# UnitBehavior - MoveTo::readData
35 # ComponentUpdate
05 00 # Component ID
# Command
01 # CreateAction2 Behavior::doLocalAction
01 # Unk
01 # ActionID

# MoveTo::readData
ff
7f 08 00 00
7f 08 00 00
#00 01 00 00
#00 00 01 00

02 00 00 00 00 # Synch
06 # End
```

```
# UnitBehavior - WarpTo::readData
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
11 # ActionID

# WarpTo::readData
01
# Values for Town
00 F0 00 00
00 F0 00 00
00 87 00 00

02 00 7e 04 00 # Synch
06 # End
```

```
# UnitBehavior - Ressurect
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
30 # ActionID

01 # FaceTarget::readInit

02 00 00 00 00 # Synch
06 # End    
```

```
# UnitBehavior - Spawn
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
04 # ActionID

# Spawn::readData
01
02 00 00 00
03 00 00 00
04 00 00 00
02 00 

02 00 7e 04 00 # Synch
06 # End    
```

```
# UnitBehavior - FollowClient
35 # ComponentUpdate
05 00 # Component ID
# Command
64
# 00 reset position (clear client control)
# 01 follow client 
01

02 00 00 00 00 # Synch
06 # End
```

```
# UnitBehavior - Kill
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
FE # ActionID

# FaceTarget::readInit
00

02 00 00 00 00 # Synch
06 # End
```

```
# UnitBehavior - Die
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
FF # ActionID

# FaceTarget::readInit
00

02 00 00 00 00 # Synch
06 # End   
```

```
# UnitBehavior - PlayAnimation
35 # ComponentUpdate
05 00 # Component ID
# Command
04 # CreateAction1
20 # ActionID

01

01 00 00 00
01 00 00 00
01 00 00 00
01 00 00 00

02 00 00 00 00 # Synch
06 # End   
```

```
# ClientEntityManager::processInterval
# This seems to log messages about the pathManager budget
# I don't think this is meant to sync each tick
0D # ID

# ClientEntityManager::processInterval
01 00 00 00
01 00 00 00
01 00 00 00

# PathManager::ReadBudget
# These values are syncing the path budget
FF FF FF FF 
FF FF # Per Update (idk)
FF FF # Per Path (idk)

06 #end
```