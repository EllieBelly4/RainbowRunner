from idc import *
from ida_bytes import *

def readTree3(addr):
    first_child_ptr = get_wide_dword(addr + 0x04)

    readTreeChildren(first_child_ptr + 4, 0)


def readTreeChildren(addr, parent_addr, depth=0):
    # if depth > 4:
    #     return

    key_ptr = addr + 12
    key_root_ptr = get_wide_dword(key_ptr)

    output_string = f"{hex(addr)} {get_wide_byte(addr + 0x14)} - {get_wide_byte(addr + 0x30)}"

    has_key = False

    if key_root_ptr != 0xBAADF00D:
        has_key = True
        type_name_ptr = get_wide_dword(key_root_ptr + 0x10)
        str_len = get_wide_dword(type_name_ptr)

        type_name = ""

        for i in range(str_len):
            type_name += chr(get_wide_byte(type_name_ptr + 0x04 + i))

        output_string += " " + type_name

    val_ptr = addr+16

    if val_ptr != 0xBAADF00D:
        val = get_wide_byte(val_ptr)

        output_string += " Val: " + str(val)

    if has_key:
        print(output_string)

    if get_wide_byte(addr + 0x15) == 1:
        return

    child_a_ptr = get_wide_dword(addr + 0x00)

    if child_a_ptr != parent_addr and child_a_ptr != 0xBAADF00D:
        readTreeChildren(child_a_ptr, addr, depth + 1)

    child_b_ptr = get_wide_dword(addr + 0x08)

    if child_b_ptr != parent_addr and child_b_ptr != 0xBAADF00D:
        readTreeChildren(child_b_ptr, addr, depth + 1)

# TreeEntry *v4; // edx
# TreeEntry *v5; // ecx
# TreeEntry *v6; // ecx
#
# v4 = *(TreeEntry **)(a2 + 4);
# v5 = (TreeEntry *)v4->field_4;
# while ( !v5->field_15 )
# {
#   if ( v5->key >= *a3 )
#   {
#     v4 = v5;
#     v5 = (TreeEntry *)v5->field_0;
#   }
#   else
#   {
#     v5 = (TreeEntry *)v5->field_8;
#   }
# }
# v6 = *(TreeEntry **)(a2 + 4);
# if ( v4 == v6 || *a3 < v4->key )
#   *result = v6;
# else
#   *result = v4;
# return result;
