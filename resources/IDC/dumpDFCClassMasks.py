from idc import *
import json

global_var_ptrs = [
    0x009306E0,
    0x009306E4,
    0x009306E8,
    0x009306F0,
    0x009306F4,
    0x009306F8,
    0x009306FC,
    0x00930700,
    0x00930710,
    0x00930714,
    0x00930718,
    0x0093071C,
    0x00930720,
    0x00930724,
    0x0093072C,
    0x00930730,
    0x00930734,
    0x00930738,
    0x0093073C,
    0x00930740,
    0x00930744,
    0x00930748,
    0x0093074C,
    0x00930750,
    0x00930754,
    0x00930758,
    0x0093075C,
    0x00930760,
    0x00930764,
    0x0093076C,
    0x00930770,
    0x00930774,
    0x00930778,
    0x0093077C,
    0x00930780,
    0x00930784,
    0x00930788,
    0x0093078C,
    0x00930790,
    0x00930794,
    0x00930798,
    0x0093079C,
    0x009307A0,
    0x009307A4,
    0x009307A8,
    0x009307AC,
    0x009307B0,
    0x009307B4,
    0x009307B8,
    0x009307BC,
    0x009307C0,
    0x009307C4,
    0x009307C8,
    0x009307CC,
    0x009307D0,
    0x009307D4,
    0x009307D8,
    0x009307DC,
    0x009307E0,
    0x009307E4,
    0x009307E8,
    0x009307EC,
    0x009307F0,
    0x009307F4,
    0x009307F8,
    0x009307FC,
    0x00930800,
    0x00930804,
    0x00930808,
    0x0093080C,
    0x00930810,
    0x00930814,
    0x00930818,
    0x0093081C,
    0x00930820,
    0x00930824,
    0x00930828,
    0x0093082C,
    0x00930830,
    0x00930834,
    0x00930838,
    0x0093083C,
    0x00930840,
    0x00930844,
    0x00930848,
    0x0093084C,
    0x00930850,
    0x00930854,
    0x00930858,
    0x0093085C,
    0x00930860,
    0x00930864,
    0x00930868,
    0x00930870,
    0x00930874,
    0x00930878,
    0x0093087C,
    0x00930880,
    0x00930884,
    0x00930888,
    0x0093088C,
    0x00930890,
    0x00930894,
    0x00930898,
    0x0093089C,
    0x009308A0,
    0x009308A8,
    0x009308AC,
    0x009308B0,
    0x009308B4,
    0x009308B8,
    0x009308BC,
    0x009308C0,
    0x009308C4,
    0x009308C8,
    0x009308CC,
    0x009308D0,
    0x009308D4,
    0x009308D8,
    0x009308DC,
    0x009308E0,
    0x009308E4,
    0x009308E8,
    0x009308EC,
    0x009308F0,
    0x009308F4,
    0x009308F8,
    0x009308FC,
    0x00930908,
    0x0093090C,
    0x00930910,
    0x00930914,
    0x00930918,
    0x0093091C,
    0x00930920,
    0x00930924,
    0x00930928,
    0x0093092C,
    0x00930930,
    0x00930934,
    0x00930938,
    0x0093093C,
    0x00930940,
    0x00930944,
    0x00930948,
    0x0093094C,
    0x00930950,
    0x00930954,
    0x00930958,
    0x0093095C,
    0x00930960,
    0x00930964,
    0x00930968,
    0x0093096C,
    0x00930970,
    0x00930974,
    0x00930978,
    0x0093097C,
    0x00930980,
    0x00930984,
    0x00930988,
    0x0093098C,
    0x00930990,
    0x00930994,
    0x00930998,
    0x0093099C,
    0x009309A0,
    0x009309A4,
    0x009309A8,
    0x009309AC,
    0x009309B0,
    0x009309B4,
    0x009309B8,
    0x009309BC,
    0x009309C0,
    0x009309C4,
    0x009309C8,
    0x009309CC,
    0x009309D0,
    0x009309D4,
    0x009309D8,
    0x009309DC,
    0x009309E0,
    0x009309E4,
    0x009309E8,
    0x009309EC,
    0x009309F0,
    0x009309F4,
    0x009309F8,
    0x009309FC,
    0x00930A00,
    0x00930A04,
    0x00930A08,
    0x00930A0C,
    0x00930A10,
    0x00930A14,
    0x00930A18,
    0x00930A1C,
    0x00930A20,
    0x00930A24,
    0x00930A28,
    0x00930A2C,
    0x00930A30,
    0x00930A34,
    0x00930A38,
    0x00930A3C,
    0x00930A40,
    0x00930A44,
    0x00930A48,
    0x00930A4C,
    0x00930A50,
    0x00930A54,
    0x00930A58,
    0x00930A5C,
    0x00930A60,
    0x00930A64,
    0x00930A68,
    0x00930A6C,
    0x00930A70,
    0x00930A74,
    0x00930A78,
    0x00930A7C,
    0x00930A80,
    0x00930A84,
    0x00930A88,
    0x00930A8C,
    0x00930A90,
    0x00930A94,
    0x00930A98,
    0x00930A9C,
    0x00930AA0,
    0x00930AA4,
    0x00930AA8,
    0x00930AAC,
    0x00930AB0,
    0x00930AB4,
    0x00930AB8,
    0x00930ABC,
    0x00930AC0,
    0x00930AC4,
    0x00930AC8,
    0x00930ACC,
    0x00930AD0,
    0x00930AD4,
    0x00930AD8,
    0x00930ADC,
    0x00930AE0,
    0x00930AE4,
    0x00930AE8,
    0x00930AEC,
    0x00930AF0,
    0x00930AF4,
    0x00930AF8,
    0x00930AFC,
    0x00930B00,
    0x00930B04,
    0x00930B08,
    0x00930B0C,
    0x00930B10,
    0x00930B14,
    0x00930B18,
    0x00930B1C,
    0x00930B20,
    0x00930B24,
    0x00930B28,
    0x00930B2C,
    0x00930B30,
    0x00930B34,
    0x00930B38,
    0x00930B3C,
    0x00930B40,
    0x00930B44,
    0x00930B48,
    0x00930B4C,
    0x00930B50,
    0x00930B54,
    0x00930B58,
    0x00930B5C,
    0x00930B60,
    0x00930B64,
    0x00930B68,
    0x00930B6C,
    0x00930B70,
    0x00930B74,
    0x00930B78,
    0x00930B7C,
    0x00930B80,
    0x00930B84,
    0x00930B88,
    0x00930B8C,
    0x00930B90,
    0x00930B94,
    0x00930B98,
    0x00930B9C,
    0x00930BA0,
    0x00930BA4,
    0x00930BA8,
    0x00930BAC,
    0x00930BB0,
    0x00930BB4,
    0x00930BB8,
    0x00930BBC,
    0x00930BC0,
    0x00930BC4,
    0x00930BC8,
    0x00930BCC,
    0x00930BD0,
    0x00930BD4,
    0x00930BD8,
    0x00930BDC,
    0x00930BE0,
    0x00930BE4,
    0x00930BE8,
    0x00930BEC,
    0x00930BF0,
    0x00930BF4,
    0x00930BF8,
    0x00930BFC,
    0x00930C00,
    0x00930C04,
    0x00930C08,
    0x00930C0C,
    0x00930C10,
    0x00930C14,
    0x00930C18,
    0x00930C1C,
    0x00930C20,
    0x00930C24,
    0x00930C28,
    0x00930C2C,
    0x00930C30,
    0x00930C34,
    0x00930C38,
    0x00930C3C,
    0x00930C40,
    0x00930C44,
    0x00930C48,
    0x00930C4C,
    0x00930C50,
    0x00930C54,
    0x00930C58,
    0x00930C5C,
    0x00930C60,
    0x00930C64,
    0x00930C68,
    0x00930C6C,
    0x00930C70,
    0x00930C74,
    0x00930C78,
    0x00930C7C,
    0x00930C80,
    0x00930C84,
    0x00930C88,
    0x00930C8C,
    0x00930C90,
    0x00930C94,
    0x00930C98,
    0x00930C9C,
    0x00930CA0,
    0x00930CA4,
    0x00930CA8,
    0x00930CAC,
    0x00930CB0,
    0x00930CB4,
    0x00930CB8,
    0x00930CBC,
    0x00930CC0,
    0x00930CC4,
    0x00930CC8,
    0x00930CCC,
    0x00930CD0,
    0x00930CD4,
    0x00930CD8,
    0x00930CDC,
    0x00930CE0,
    0x00930CE4,
    0x00930CEC,
    0x00930CF0,
    0x00930CF4,
    0x00930CF8,
    0x00930CFC,
    0x00930D00,
    0x00930D04,
    0x00930D08,
    0x00930D0C,
    0x00930D10,
    0x00930D14,
    0x00930D18,
    0x00930D1C,
    0x00930D20,
    0x00930D24,
    0x00930D28,
    0x00930D2C,
    0x00930D34,
    0x00930D38,
    0x00930D40,
    0x00930D44,
    0x00930D48,
    0x00930D4C,
    0x00930D50,
    0x00930D54,
    0x00930D58,
    0x00930D5C,
    0x00930D60,
    0x00930D64,
    0x00930D6C,
    0x00930D70,
    0x00930D74,
    0x00930D78,
    0x00930D7C,
    0x00930D80,
    0x00930DA0,
    0x00930DA4,
    0x00930DA8,
    0x00930DAC,
    0x00930DB0,
    0x00930DC4,
    0x00930DC8,
    0x00930DCC,
    0x00930DD0,
    0x00930DD4,
    0x00930DD8,
    0x00930DDC,
    0x00930DE0,
    0x00930DE4,
    0x00930DE8,
    0x00930DEC,
    0x00930DF0,
    0x00930DF8,
    0x00930DFC,
    0x00930E00,
    0x00930E04,
    0x00930E08,
    0x00930E0C,
    0x00930E10,
    0x00930E14,
    0x00930E18,
    0x00930E1C,
    0x00930E20,
    0x00930E24,
    0x00930E28,
    0x00930E2C,
    0x00930E30,
    0x00930E34,
    0x00930E38,
    0x00930E3C,
    0x00930E40,
    0x00930E4C,
    0x00930E50,
    0x00930E54,
    0x00930E58,
    0x00930E5C,
    0x00930E60,
    0x00930E64,
    0x00930E68,
    0x00930E6C,
    0x00930E70,
    0x00930E78,
    0x00930E7C,
    0x00930E80,
    0x00930E84,
    0x00930E88,
    0x00930E8C,
    0x00930E90,
    0x00930E94,
    0x00930E98,
    0x00930E9C,
    0x00930EA0,
    0x00930EA4,
    0x00930EA8,
    0x00930EAC,
    0x00930EB0,
    0x00930EB4,
    0x00930EB8,
    0x00930EBC,
    0x00930EC0,
    0x00930EC4,
    0x00930EC8,
    0x00930ECC,
    0x00930ED0,
    0x00930ED4,
    0x00930ED8,
    0x00930EE0,
    0x00930EE4,
    0x00930EE8,
    0x00930EEC,
    0x00930EF0,
    0x00930EF4,
    0x00930EF8,
    0x00930EFC,
    0x00930F00,
    0x00930F04,
    0x00930F08,
    0x00930F0C,
    0x00930F10,
    0x00930F14,
    0x00930F18,
    0x00930F1C,
    0x00930F20,
    0x00930F24,
    0x00930F28,
    0x00930F2C,
    0x00930F30,
    0x00930F34,
    0x00930F38,
    0x00930F3C,
    0x00930F40,
    0x00930F44,
    0x00930F4C,
    0x00930F50,
    0x00930F54,
    0x00930F58,
    0x00930F5C,
    0x00930F60,
    0x00930F64,
    0x00930F68,
    0x00930F6C,
    0x00930F70,
    0x00930F74,
    0x00930F78,
    0x00930F7C,
    0x00930F80,
    0x00930F84,
    0x00930F88,
    0x00930F8C,
    0x00930F90,
    0x00930F94,
    0x00930F98,
    0x00930F9C,
    0x00930FA0,
    0x00930FA4,
    0x00930FA8,
    0x00930FAC,
    0x00930FB0,
    0x00930FB4,
    0x00930FB8,
    0x00930FBC,
    0x00930FC0,
    0x00930FC4,
    0x00930FC8,
    0x00930FCC,
    0x00930FD0,
    0x00930FD4,
    0x00930FD8,
    0x00930FDC,
    0x00930FE0,
    0x00930FE4,
    0x00930FE8,
    0x00930FEC,
    0x00930FF4,
    0x00930FF8,
    0x00930FFC,
    0x00931000,
    0x00931004,
    0x00931008,
    0x0093100C,
    0x00931010,
    0x00931014,
    0x00931018,
    0x0093101C,
    0x00931020,
    0x00931024,
    0x00931028,
    0x00931030,
    0x00931038,
    0x0093103C,
    0x00931040,
    0x00931054,
    0x00931058,
    0x0093105C,
    0x00931060,
    0x00931064,
    0x00931068,
]


def read_dr_string(ptr_drstring):
    string_length = get_wide_dword(ptr_drstring)
    res_str = ""

    for i in range(string_length):
        res_str += chr(get_wide_byte(ptr_drstring + 4 + i))

    return res_str


def get_dfc_class_data(ptr_global_var):
    print(f"getting {hex(ptr_global_var)}")

    ptr_dfclass = get_wide_dword(ptr_global_var)

    ptr_mask = ptr_dfclass + 0x40
    mask0 = get_wide_dword(ptr_mask)
    mask1 = get_wide_dword(ptr_mask + 4)

    ptr_name = get_wide_dword(ptr_dfclass + 0x10)
    name = "unknown"
    if ptr_name != 0:
        name = read_dr_string(ptr_name)

    something_class_type = get_wide_dword(ptr_dfclass + 0x38)
    class_type = get_wide_dword(ptr_dfclass + 0x98)

    return {
        "mask0": mask0,
        "mask1": mask1,
        "name": name,
        "class_type": class_type,
        "something_class_type": something_class_type
    }


def get_all_dfc_class_data(ptrs):
    results = {}

    for ptr_global_var in ptrs:
        if get_wide_dword(ptr_global_var) == 0:
            continue

        result = get_dfc_class_data(ptr_global_var)
        results[result["name"]] = result

    return results


results = get_all_dfc_class_data(global_var_ptrs)
# results = get_dfc_class_data(0x00930CC4)

# Write results JSON to file
with open("dfc_class_data.json", "w") as f:
    json.dump(results, f, indent=2)
