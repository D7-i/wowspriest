package shaman

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const AncestralAwakeningHealMultiplier = 0.3

func (shaman *Shaman) applyAncestralAwakening() {
	if !shaman.HasRune(proto.ShamanRune_RuneFeetAncestralAwakening) {
		return
	}

	shaman.AncestralAwakening = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: int32(proto.ShamanRune_RuneFeetAncestralAwakening)},
		SpellSchool: core.SpellSchoolNature,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellHealing,
		Flags:       SpellFlagShaman | core.SpellFlagHelpful | core.SpellFlagAPL,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealHealing(sim, target, shaman.ancestralHealingAmount*(1+shaman.purificationHealingModifier()), spell.OutcomeHealing)
		},
	})
}
