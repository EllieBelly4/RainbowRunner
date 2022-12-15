"SessionID" is the number that increments when an action is sent... sometimes.
It is incremented here:
`.text:00520410     void_UnitBehavior__SyncClient_class_Fixed32_bool_1 proc near`

This can happen when the client has not received an update in more than 10 ticks or 
whenever the client sends StopFollowClient