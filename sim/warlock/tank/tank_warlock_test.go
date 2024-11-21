package tank

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterTankWarlock()
}

func TestAffliction(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 4,
			Level: 60,
			Race:  proto.Race_RaceOrc,

			Talents:     Phase4AffTalents,
			GearSet:     core.GetGearSet("../../../ui/tank_warlock/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/tank_warlock/apls", "p4_destro_aff_tank"),
			Buffs:       core.FullBuffsPhase4,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Affliction Warlock", SpecOptions: DefaultAfflictionWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestDemonology(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 4,
			Level: 60,
			Race:  proto.Race_RaceOrc,

			Talents:     Phase4DemoTalents,
			GearSet:     core.GetGearSet("../../../ui/tank_warlock/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/tank_warlock/apls", "p4_demo_tank"),
			Buffs:       core.FullBuffsPhase4,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Demonology Warlock", SpecOptions: DefaultDemonologyWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestDestruction(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 4,
			Level: 60,
			Race:  proto.Race_RaceOrc,

			Talents:     Phase4DestroTalents,
			GearSet:     core.GetGearSet("../../../ui/tank_warlock/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/tank_warlock/apls", "p4_destro_aff_tank"),
			Buffs:       core.FullBuffsPhase4,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Destruction Warlock", SpecOptions: DefaultDestroWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var Phase1AfflictionTalents = "05002-005"
var Phase1DestructionTalents = "-03-0550201"

var Phase2DemonologyTalents = "-2050033112501251"
var Phase2DestructionTalents = "-035-05500050025001"

var Phase3DestructionTalents = "05-03-505020500050515"

var Phase4AffTalents = "5500253011201002-03-50502051002001"
var Phase4DemoTalents = "-205004015250105-50500050005001"
var Phase4DestroTalents = "45002400102-03-505020510050115"

var DefaultDestroWarlock = &proto.Player_TankWarlock{
	TankWarlock: &proto.TankWarlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_FelArmor,
			Summon:      proto.WarlockOptions_Imp,
			WeaponImbue: proto.WarlockOptions_NoWeaponImbue,
		},
	},
}

var DefaultAfflictionWarlock = &proto.Player_TankWarlock{
	TankWarlock: &proto.TankWarlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_FelArmor,
			Summon:      proto.WarlockOptions_Imp,
			WeaponImbue: proto.WarlockOptions_NoWeaponImbue,
		},
	},
}

var DefaultDemonologyWarlock = &proto.Player_TankWarlock{
	TankWarlock: &proto.TankWarlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_FelArmor,
			Summon:      proto.WarlockOptions_Felguard,
			WeaponImbue: proto.WarlockOptions_Firestone,
		},
	},
}

var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir: proto.AgilityElixir_ElixirOfLesserAgility,
		DefaultPotion: proto.Potions_ManaPotion,
		FirePowerBuff: proto.FirePowerBuff_ElixirOfFirepower,
		Food:          proto.Food_FoodSmokedSagefish,
		MainHandImbue: proto.WeaponImbue_BlackfathomManaOil,
		StrengthBuff:  proto.StrengthBuff_ElixirOfOgresStrength,
	},
}

var Phase2Consumes = core.ConsumesCombo{
	Label: "P2-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:  proto.Potions_ManaPotion,
		FirePowerBuff:  proto.FirePowerBuff_ElixirOfFirepower,
		Food:           proto.Food_FoodSagefishDelight,
		MainHandImbue:  proto.WeaponImbue_LesserWizardOil,
		SpellPowerBuff: proto.SpellPowerBuff_LesserArcaneElixir,
	},
}

var Phase3Consumes = core.ConsumesCombo{
	Label: "P3-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:   proto.Potions_SuperiorManaPotion,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfFirepower,
		ShadowPowerBuff: proto.ShadowPowerBuff_ElixirOfShadowPower,
		Food:            proto.Food_FoodTenderWolfSteak,
		MainHandImbue:   proto.WeaponImbue_LesserWizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var Phase4Consumes = core.ConsumesCombo{
	Label: "P4-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:   proto.Potions_MajorManaPotion,
		Flask:           proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfGreaterFirepower,
		ShadowPowerBuff: proto.ShadowPowerBuff_ElixirOfShadowPower,
		Food:            proto.Food_FoodTenderWolfSteak,
		MainHandImbue:   proto.WeaponImbue_WizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeDagger,
	},
	HandTypes: []proto.HandType{
		proto.HandType_HandTypeOffHand,
	},
	ArmorType: proto.ArmorType_ArmorTypeCloth,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeWand,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpellPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}
