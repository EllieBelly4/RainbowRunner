package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/serverconfig"
	"RainbowRunner/internal/types/drobjecttypes"
	"fmt"
	lua2 "github.com/yuin/gopher-lua"
)

func registerLuaObjectHelpers(s *lua2.LState) {
	mt := s.NewTypeMetatable("ObjectHelpers")
	s.SetGlobal("ObjectHelpers", mt)
	s.SetFuncs(mt, map[string]lua2.LGFunction{
		"loadAvatar": func(state *lua2.LState) int {
			avatar := LoadAvatar()
			avatarLua := avatar.ToLua(state)
			state.Push(avatarLua)
			return 1
		},
	})
}

func LoadAvatar() *Avatar {
	avatar := NewAvatar("avatar.classes.FighterFemale")
	avatar.GCLabel = "Avatar Name"
	avatar.Properties = []GCObjectProperty{
		Uint32Prop("Hair", 0x01),
		Uint32Prop("HairColor", 0x00),
		Uint32Prop("Face", 0),
		Uint32Prop("FaceFeature", 0),
		Uint32Prop("Skin", 0x01),
		Uint32Prop("Level", 50),
	}

	//metrics := NewAvatarMetrics(0xFE34BE34, "EllieMetrics")

	modifiers := NewModifiers("Modifiers")
	modifiers.GCLabel = "Mod Name"
	modifiers.Properties = []GCObjectProperty{
		Uint32Prop("IDGenerator", 0x01),
	}

	manipulators := NewManipulators("Manipulators")
	manipulators.GCLabel = "ManipulateMe"

	dialogManager := NewGCObject("DialogManager")
	dialogManager.GCLabel = "EllieDialogManager"

	questManager := NewQuestManager()
	questManager.GCLabel = "EllieQuestManager"

	//animationList := objects.NewGCObject("AnimationList")
	//animationList.Name = "EllieAnimations"

	avatarSkills := NewSkills("avatar.base.skills")
	avatarSkills.GCLabel = "EllieSkills"

	//skillSlot := objects.NewSkillSlot("skillslot")
	//skillSlot.GCLabel = "EllieSkillSlot"
	//skillSlot.SlotID = 0x64
	//skillSlot.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotID", 0x64),
	//}
	//avatarSkills.AddChild(skillSlot)

	skillsToAdd := []struct {
		Name  string
		Level byte
		Unk0  uint32
	}{
		{"skills.generic.Stomp", 1, 1},
		{"skills.generic.Sprint", 1, 2},
		{"skills.generic.Butcher", 1, 3},
		{"skills.generic.Blight", 1, 5},
		{"skills.generic.Charge", 1, 6},
		{"skills.generic.Cleave", 1, 7},
		{"skills.generic.IceBolt", 1, 9},
		{"skills.generic.IceShot", 1, 10},
		{"skills.generic.ManaShield", 1, 11},
		{"skills.generic.FearShot", 1, 10000},
	}

	for i, s := range skillsToAdd {
		skill := NewActiveSkill(s.Name)
		skill.Level = s.Level
		skill.GCLabel = s.Name

		skill.Properties = []GCObjectProperty{
			//objects.Uint32Prop("Level", s.Level),
		}

		hotbarSlot := i + 0x64

		if hotbarSlot >= 0x6D {
			hotbarSlot = i + 1
		}

		if hotbarSlot >= 0x64 {
			manipulators.AddChild(skill)
		}

		avatarSkills.AddSkill(skill, i+1, hotbarSlot)
	}

	//skillSlot := objects.NewComponent("skillslot", "skillslot")
	//skillSlot.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotID", 0x01),
	//}
	//avatarSkills.AddChild(skillSlot)

	//avatarDesc := objects.NewGCObject("AvatarDesc")
	//avatarDesc.GCType = "avatar.classes.fighterfemale.description"
	//avatarDesc.Name = "EllieAvatarDesc"

	//avatarEquipment := objects.NewGCObject("Equipment")
	//avatarEquipment.GCType = "avatar.base.Equipment"
	//avatarEquipment.GCLabel = "EllieEquipment"

	avatarEquipment := NewEquipmentInventory("avatar.base.Equipment", avatar)
	avatarEquipment.GCLabel = "EllieEquipment"

	//.text:0058E550     ; struct DFCClass *__thiscall Armor::getClass(Armor *__hidden this)
	//.text:0058E550 000 mov     eax, ?Class@Armor@@2PAVDFCClass@@A ; DFCClass * Armor::Class

	//weapon := objects.NewGCObject("MeleeWeapon")
	//weapon.GCType = "1HSwordMythicPAL.1HSwordMythic6"
	//weapon.GCLabel = "EllieWeapon"

	//weaponDesc := objects.NewGCObject("MeleeWeaponDesc")
	//weaponDesc.GCType = "1HMace1PAL.1HMace1-1.Description"
	//weaponDesc.GCLabel = "EllieWeaponDesc"
	//weaponDesc.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotType", uint32(objects.EquipmentSlotWeapon)),
	//}

	//TODO finish
	//weapon.Properties = []objects.GCObjectProperty{
	////	objects.Uint32Prop("ItemDesc.SlotType", 0x0a),
	//	objects.Uint32Prop("EquipmentSlot", uint32(EquipmentSlotWeapon)),
	//}

	//armor.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("Slot", uint32(EquipmentSlotTorso)),
	//}

	manipulator := NewGCObject("Manipulator")
	//manipulator.GCType = "base.MeleeUnit.Manipulators.PrimaryWeapon"
	//objects.Entities.RegisterAll(conn, manipulator)

	unitContainer := NewUnitContainer(manipulator, "EllieUnitContainer", avatar)
	//unitContainer.GCType = "unitcontainer"
	//unitContainer.Name = "EllieUnitContainer"

	baseInventory := NewInventory("avatar.base.Inventory", 11)
	baseInventory.GCLabel = "EllieBaseInventory"

	bankInventory := NewInventory("avatar.base.Bank", 12)
	bankInventory.GCLabel = "EllieBankInventory"

	tradeInventory := NewInventory("avatar.base.TradeInventory", 13)
	tradeInventory.GCLabel = "EllieTradeInventory"

	// Items in inventories
	//randomItem := objects.NewEquipment("PlateMythicPAL.PlateMythicBoots1", "PlateMythicPAL.PlateMythicBoots1.Mod1", objects.ItemArmour, types.EquipmentSlotFoot)
	//baseInventory.AddChild(randomItem)

	//r := rand.New(rand.NewSource(time.Now().Unix()))

	AddEquipment(avatarEquipment, manipulators,
		"PlateArmor3PAL.PlateArmor3-7",
		"PlateBoots3PAL.PlateBoots3-7",
		"PlateHelm3PAL.PlateHelm3-7",
		"PlateGloves3PAL.PlateGloves3-7",
		"CrystalMythicPAL.CrystalMythicShield1",
	)

	unitBehaviour := NewUnitBehavior("avatar.base.UnitBehavior")
	unitBehaviour.UnitBehaviorUnk1 = 0x01
	unitBehaviour.UnitBehaviorUnk2 = 0x01
	unitBehaviour.UnitMoverFlags |= 0x08
	unitBehaviour.UnitMoverUnk0 = 0x01
	unitBehaviour.UnitBehaviorTicksSinceLastUpdate = 0x10

	unitBehaviour.IsUnderClientControl = false

	unitBehaviour.GCLabel = "EllieBehaviour"

	unitContainer.AddChild(baseInventory)
	unitContainer.AddChild(bankInventory)
	unitContainer.AddChild(tradeInventory)

	//avatar.AddChild(visual)
	//avatar.AddChild(rpgSettings)
	avatar.AddChild(avatarEquipment)
	avatar.AddChild(avatarSkills)
	avatar.AddChild(unitContainer)
	avatar.AddChild(unitBehaviour)
	avatar.AddChild(modifiers)
	avatar.AddChild(manipulators)
	//avatar.AddChild(metrics)
	avatar.AddChild(dialogManager)
	avatar.AddChild(questManager)
	//avatar.AddChild(avatarDesc)

	return avatar
}

func AddEquipment(equipment drobjecttypes.DRObject, manipulators *Manipulators, armour string, boots string, helm string, gloves string, shield string) {
	randomArmour := AddRandomEquipment(database.Armours, ItemArmour)

	if randomArmour != nil {
		equipment.AddChild(randomArmour)
		manipulators.AddChild(randomArmour)
	}

	randomBoots := AddRandomEquipment(database.Boots, ItemArmour)

	if randomBoots != nil {
		equipment.AddChild(randomBoots)
		manipulators.AddChild(randomBoots)
	}

	randomHelm := AddRandomEquipment(database.Helmets, ItemArmour)

	if randomBoots != nil {
		equipment.AddChild(randomHelm)
		manipulators.AddChild(randomHelm)
	}

	randomGloves := AddRandomEquipment(database.Gloves, ItemArmour)

	if randomGloves != nil {
		equipment.AddChild(randomGloves)
		manipulators.AddChild(randomGloves)
	}

	randomWeapon := AddRandomEquipment(database.MeleeWeapons, ItemMeleeWeapon)
	if randomWeapon != nil {
		equipment.AddChild(randomWeapon)
		manipulators.AddChild(randomWeapon)
	}

	if serverconfig.Config.Logging.LogRandomEquipment {
		randomArmourName := "None"
		if randomArmour != nil {
			randomArmourName = randomArmour.GetGCType()
		}

		randomBootsName := "None"
		if randomBoots != nil {
			randomBootsName = randomBoots.GetGCType()
		}

		randomHelmName := "None"
		if randomHelm != nil {
			randomHelmName = randomHelm.GetGCType()
		}

		randomGlovesName := "None"
		if randomGloves != nil {
			randomGlovesName = randomGloves.GetGCType()
		}

		randomWeaponName := "None"
		if randomWeapon != nil {
			randomWeaponName = randomWeapon.GetGCType()
		}

		fmt.Printf(`Random equipment for today is:
Helm: %s
Armour: %s
Gloves: %s
Boots: %s
Weapon: %s
`, randomHelmName, randomArmourName, randomGlovesName, randomBootsName, randomWeaponName)
	}
}
