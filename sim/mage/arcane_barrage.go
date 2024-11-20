package mage

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (mage *Mage) registerArcaneBarrageSpell() {
	if !mage.HasRune(proto.MageRune_RuneCloakArcaneBarrage) {
		return
	}

	baseDamageLow := mage.baseRuneAbilityDamage() * 3.58
	baseDamageHigh := mage.baseRuneAbilityDamage() * 4.38
	damageCoef := 0.429
	manaCost := 0.08
	cooldown := time.Second * 3

	mage.ArcaneBarrage = mage.RegisterSpell(core.SpellConfig{
		SpellCode:    SpellCode_MageArcaneBarrage,
		ActionID:     core.ActionID{SpellID: int32(proto.MageRune_RuneCloakArcaneBarrage)},
		SpellSchool:  core.SpellSchoolArcane,
		DefenseType:  core.DefenseTypeMagic,
		ProcMask:     core.ProcMaskSpellDamage,
		Flags:        SpellFlagMage | core.SpellFlagAPL,
		MissileSpeed: 24,

		ManaCost: core.ManaCostOptions{
			BaseCost: manaCost,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    mage.NewTimer(),
				Duration: cooldown,
			},
		},

		BonusCoefficient: damageCoef,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcDamage(sim, target, sim.Roll(baseDamageLow, baseDamageHigh), spell.OutcomeMagicHitAndCrit)

			spell.WaitTravelTime(sim, func(sim *core.Simulation) {
				spell.DealDamage(sim, result)
			})
		},
	})
}
