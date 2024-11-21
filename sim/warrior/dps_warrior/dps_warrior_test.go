package dpswarrior

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common" // imported to get item effects included.
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterDpsWarrior()
}

func TestDualWieldWarrior(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassWarrior,
			Phase:      5,
			Level:      60,
			Race:       proto.Race_RaceOrc,
			OtherRaces: []proto.Race{proto.Race_RaceHuman},

			Talents:     P4FuryTalents,
			GearSet:     core.GetGearSet("../../../ui/warrior/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/warrior/apls", "phase_5_dw"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Fury", SpecOptions: PlayerOptionsFury},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestTwoHandedWarrior(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassWarrior,
			Phase:      5,
			Level:      60,
			Race:       proto.Race_RaceOrc,
			OtherRaces: []proto.Race{proto.Race_RaceHuman},

			Talents:     P4FuryTalents,
			GearSet:     core.GetGearSet("../../../ui/warrior/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/warrior/apls", "phase_5_2h"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Arms", SpecOptions: PlayerOptionsArms},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var P2ArmsTalents = "303050213525100001"
var P2FuryTalents = "-05050005405010051"
var P3ArmsTalents = "303050213520105001-0505"
var P4FuryTalents = "20305020302-05050005525010051"

var Phase4Consumes = core.ConsumesCombo{
	Label: "P4-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		DefaultPotion:     proto.Potions_MightyRagePotion,
		DragonBreathChili: true,
		Food:              proto.Food_FoodSmokedDesertDumpling,
		MainHandImbue:     proto.WeaponImbue_Windfury,
		OffHandImbue:      proto.WeaponImbue_ElementalSharpeningStone,
		StrengthBuff:      proto.StrengthBuff_JujuPower,
	},
}

var PlayerOptionsArms = &proto.Player_Warrior{
	Warrior: &proto.Warrior{
		Options: warriorOptions,
	},
}

var PlayerOptionsFury = &proto.Player_Warrior{
	Warrior: &proto.Warrior{
		Options: warriorOptions,
	},
}

var warriorOptions = &proto.Warrior_Options{
	StartingRage: 50,
	Shout:        proto.WarriorShout_WarriorShoutBattle,
}

var ItemFilters = core.ItemFilter{
	ArmorType: proto.ArmorType_ArmorTypePlate,

	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeAxe,
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeFist,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatStrength,
	proto.Stat_StatAgility,
	proto.Stat_StatAttackPower,
	proto.Stat_StatMeleeCrit,
	proto.Stat_StatMeleeHit,
}
