package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const ShamanMoltenBlastResetChance = .10

func (shaman *Shaman) applyMoltenBlast() {
	if !shaman.HasRune(proto.ShamanRune_RuneHandsMoltenBlast) {
		return
	}

	baseDamageLow := shaman.baseRuneAbilityDamage() * .72
	baseDamageHigh := shaman.baseRuneAbilityDamage() * 1.08
	apCoef := .10
	spCoef := .14
	cooldown := time.Second * 6
	manaCost := .18
	targetCount := int32(10)

	numHits := min(targetCount, shaman.Env.GetNumTargets())
	results := make([]*core.SpellResult, numHits)

	shaman.MoltenBlast = shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: int32(proto.ShamanRune_RuneHandsMoltenBlast)},
		SpellCode:   SpellCode_ShamanMoltenBlast,
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagShaman | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: manaCost,
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    shaman.NewTimer(),
				Duration: cooldown,
			},
		},

		BonusCoefficient: spCoef,
		DamageMultiplier: 1,
		ThreatMultiplier: 2,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for idx := range results {
				// Molten Blast is a magic ability but scales off of Attack Power
				baseDamage := sim.Roll(baseDamageLow, baseDamageHigh) + apCoef*spell.MeleeAttackPower()
				results[idx] = spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
				target = sim.Environment.NextTargetUnit(target)
			}

			for _, result := range results {
				spell.DealDamage(sim, result)
			}
		},
	})
}
