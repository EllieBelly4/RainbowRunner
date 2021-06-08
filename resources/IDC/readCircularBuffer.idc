static readCircularBufferMovementUpdate(addr){
    auto next = dword(addr);
    auto previous = dword(addr + 4);
    
    print(next);
    print(previous);
     
    
//    auto i;
//    auto str = "";
//    auto len = Dword(addr);
//    
//    for(i = 0; i < len; i++){
//        str = str + Byte(addr + 4 + i);
//    }
//    
//    print(str);
}
