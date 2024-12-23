package druid

import (
	"github.com/wowsims/classic/sim/core"
)

const SwipeRanks = 5

var SwipeSpellId = [SwipeRanks + 1]int32{0, 779, 780, 769, 9754, 9908}
var SwipeBaseDamage = [SwipeRanks + 1]float64{0, 18, 25, 36, 60, 83}
var SwipeLevel = [SwipeRanks + 1]int{0, 16, 24, 34, 44, 54}

// See https://www.wowhead.com/classic/spell=436895/s03-tuning-and-overrides-passive-druid
// Modifies Threat +101%:
const SwipeThreatMultiplier = 2.0

func (druid *Druid) registerSwipeBearSpell() {
	rank := map[int32]int{
		25: 2,
		40: 3,
		50: 4,
		60: 6,
	}[druid.Level]

	level := SwipeLevel[rank]
	spellID := SwipeSpellId[rank]
	baseDamage := SwipeBaseDamage[rank]

	rageCost := 20 - float64(druid.Talents.Ferocity)
	numHits := min(3, druid.Env.GetNumTargets())
	results := make([]*core.SpellResult, numHits)

	switch druid.Ranged().ID {
	case IdolOfBrutality:
		rageCost -= 3
	}

	druid.SwipeBear = druid.RegisterSpell(Bear, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       SpellFlagOmen | core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		Rank:          rank,
		RequiredLevel: level,

		RageCost: core.RageCostOptions{
			Cost: 20 - float64(druid.Talents.Ferocity),
		},

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		DamageMultiplier: 1 + 0.1*float64(druid.Talents.SavageFury),
		ThreatMultiplier: SwipeThreatMultiplier,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for idx := range results {
				results[idx] = spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)
				target = sim.Environment.NextTargetUnit(target)
			}

			for _, result := range results {
				spell.DealDamage(sim, result)
			}
		},
	})
}
