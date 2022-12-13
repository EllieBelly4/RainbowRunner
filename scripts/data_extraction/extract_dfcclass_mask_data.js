const fs = require("fs")

const data = JSON.parse(fs.readFileSync("../../resources/Dumps/dfc_class_data.json"))
const maskNameList = [
    "unk0",
    "unk1",
    "unk2",
    "unk3",
    "unk4",
    "unk5",
    "unk6",
    "unk7",
    "unk8",
    "unk9",
    "unk10",
    "unk11",
    "unk12",
    "unk13",
    "unk14",
    "unk15",
    "unk16",
    "unk17",
    "Object?",
    "unk19",
    "Desc?",
    "StateObject?",
    "StaticObjectDesc?",
    "Visual?",
    "MountedVisual?",
    "DyeVisual?",
    "TypeSwitchVisual?",
    "RemoveNodeVisual?",
    "AnimateNodeVisual?",
    "World?",
    "unk30",
    "Entity",
    "unk32",
    "unk33",
    "unk34",
    "unk35",
    "unk36",
    "unk37",
    "unk38",
    "unk39",
    "unk40",
    "unk41",
    "unk42",
    "unk43",
    "unk44",
    "unk45",
    "unk46",
    "unk47",
    "unk48",
    "unk49",
    "unk50",
    "unk51",
    "unk52",
    "unk53",
    "unk54",
    "unk55",
    "unk56",
    "unk57",
    "unk58",
    "unk59",
    "unk60",
    "unk61",
    "unk62",
    "UnitComponent?",
]

const maskGrouped = {}

for (let i = 0; i < 64; i++) {
    const name = maskNameList[i]
    maskGrouped[name] = []
}

for (const [key, value] of Object.entries(data)) {
    const checkMask = value.mask0 == 0 ? value.mask1 : value.mask0
    const offset = value.mask0 == 0 ? 32 : 0

    for (let i = 0; i < 32; i++) {
        const name = maskNameList[i + offset]

        if (checkMask & (1 << i)) {
            maskGrouped[name].push(key)
        }
    }

}

let json = JSON.stringify(maskGrouped, null, 2);

fs.writeFileSync("../../resources/Dumps/dfc_class_mask_data.json", json)
