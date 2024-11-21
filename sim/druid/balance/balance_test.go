package balance

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common" // imported to get caster sets included. (we use spellfire here)
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterBalanceDruid()
}

func TestBalance(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassDruid,
			Phase:      5,
			Level:      60,
			Race:       proto.Race_RaceTauren,
			OtherRaces: []proto.Race{proto.Race_RaceNightElf},

			Talents:     Phase4Talents,
			GearSet:     core.GetGearSet("../../../ui/balance_druid/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/balance_druid/apls", "phase_5"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase5Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Default", SpecOptions: PlayerOptionsAdaptive},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var Phase1Talents = "50005003021"
var Phase2Talents = "5000500302541051"
var Phase3Talents = "5000550012551351--3"
var Phase4Talents = "5000550012551251--5005031"

var Phase5Consumes = core.ConsumesCombo{
	Label: "P5-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:  proto.Potions_MajorManaPotion,
		Flask:          proto.Flask_FlaskOfSupremePower,
		Food:           proto.Food_FoodNightfinSoup,
		MainHandImbue:  proto.WeaponImbue_BrilliantWizardOil,
		SpellPowerBuff: proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var PlayerOptionsAdaptive = &proto.Player_BalanceDruid{
	BalanceDruid: &proto.BalanceDruid{
		Options: &proto.BalanceDruid_Options{
			OkfUptime: 0.2,
		},
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeOffHand,
		proto.WeaponType_WeaponTypeStaff,
		proto.WeaponType_WeaponTypePolearm,
	},
	ArmorType: proto.ArmorType_ArmorTypeLeather,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeIdol,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpellPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}
