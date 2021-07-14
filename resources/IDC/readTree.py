from idc import *
from ida_bytes import *

def readTree(addr, depth = 0):

    if depth > 4:
        return

    search_addr = addr

    key = get_wide_dword(search_addr + 12)
    
    classAddr = get_wide_dword(search_addr + 16)
    vftableAddr = get_wide_dword(classAddr);

    name = get_name(vftableAddr)

    print(f"{name} {hex(key)} value {hex(vftableAddr)}")
    
    search0 = get_wide_dword(search_addr)
    
    if get_byte(search0+0x15) == 0:
        readTree(search0, depth + 1)
        
    search8 = get_wide_dword(search_addr + 0x08)
    
    if get_byte(search8 + 0x15) == 0:
        readTree(search8, depth + 1)

readTree(0x028FB8F8)