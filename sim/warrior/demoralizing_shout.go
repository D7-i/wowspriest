package warrior

import (
	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerDemoralizingShoutSpell() {
	rank := core.LevelToDebuffRank[core.DemoralizingShout][warrior.Level]
	actionId := core.DemoralizingShoutSpellId[rank]

	warrior.DemoralizingShoutAuras = warrior.NewEnemyAuraArray(func(target *core.Unit, level int32) *core.Aura {
		return core.DemoralizingShoutAura(target, warrior.Talents.BoomingVoice, warrior.Talents.ImprovedDemoralizingShout, warrior.Level)
	})

	warrior.DemoralizingShout = warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: actionId},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskEmpty,
		Flags:       core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost: 10,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ThreatMultiplier: 0.4,
		FlatThreatBonus:  0.4 * 2 * float64(core.DemoralizingShoutLevel[rank]),

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for _, aoeTarget := range sim.Encounter.TargetUnits {
				result := spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeMagicHit)
				if result.Landed() {
					warrior.DemoralizingShoutAuras.Get(aoeTarget).Activate(sim)
				}
			}
		},

		RelatedAuras: []core.AuraArray{warrior.DemoralizingShoutAuras},
	})
}
