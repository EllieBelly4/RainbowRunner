from idc import *

def readCircularBuffer(addr):
    i = 0

    search_addr = addr

    while i < 100:       
        prev_addr = get_wide_dword(search_addr)

        world_input_handler = get_wide_dword(search_addr+8)
        rotation = get_wide_dword(search_addr+12)
        posX = get_wide_dword(search_addr+16)
        posY = get_wide_dword(search_addr+20)

        print("Rot: {:x} Pos: {:x} {:x}".format(rotation, posX, posY))

        search_addr = get_wide_dword(search_addr + 4)

        if search_addr == addr:
            break

        i += 1
    
    print(f"Found {i+1}")


readCircularBuffer(0x02A361E8)