checkpoint->class_type = 0x16C
checkpoint->something_class_type = 0x400008
entity->class_type = 0x5F

// 0x5F >> 5 = 0x02
og_type_shift = dfcclass->class_type >> 5

// (0x20 * og_type_shift) = 0x40
// 0x5F - 0x40 = 0x1F
// 1 << 0x1F = 0x80000000
og_type_mask = 1 << (LOBYTE(dfcclass->class_type) - 32 * og_type_shift)

// 0x400008 + og_type_shift = 0x40000A
// 0x80000000 & 0x40000A = 0x00
(og_type_mask & *(&dfcclass_to_check->something_class_type + og_type_shift)) == og_type_mask


