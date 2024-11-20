package mage

import (
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// TODO: Classic verify Arcane Blast rune numbers
// https://www.wowhead.com/classic/news/patch-1-15-build-52124-ptr-datamining-season-of-discovery-runes-336044#news-post-336044
// https://www.wowhead.com/classic/spell=400574/arcane-blast
func (mage *Mage) registerArcaneBlastSpell() {
	if !mage.HasRune(proto.MageRune_RuneHandsArcaneBlast) {
		return
	}

	hasLivingFlameRune := mage.HasRune(proto.MageRune_RuneLegsLivingFlame)

	baseLowDamage := mage.baseRuneAbilityDamage() * 4.53
	baseHighDamage := mage.baseRuneAbilityDamage() * 5.27
	spellCoeff := .714
	castTime := time.Millisecond * 2500
	manaCost := .07

	mage.ArcaneBlastDamageMultiplier = 0.15

	additiveDamageAffectedSpells := []*core.Spell{}
	// Purposefully excluded living flame and arcane missiles ticks because we manually disable the arcane blast aura after the final tick
	affectedSpellCodes := []int32{
		SpellCode_MageArcaneBarrage, SpellCode_MageArcaneExplosion, SpellCode_MageArcaneSurge, SpellCode_MageBalefireBolt, SpellCode_MageSpellfrostBolt,
	}

	mage.ArcaneBlastAura = mage.GetOrRegisterAura(core.Aura{
		Label:     "Arcane Blast Aura",
		ActionID:  core.ActionID{SpellID: 400573},
		Duration:  time.Second * 6,
		MaxStacks: 4,
		OnInit: func(aura *core.Aura, sim *core.Simulation) {
			additiveDamageAffectedSpells = core.FilterSlice(
				core.Flatten([][]*core.Spell{
					mage.ArcaneExplosion,
					mage.ArcaneMissilesTickSpell,
					{mage.ArcaneSurge},
					{mage.SpellfrostBolt},
					{mage.BalefireBolt},
				}),
				func(spell *core.Spell) bool { return spell != nil },
			)
		},
		OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks int32, newStacks int32) {
			mage.ArcaneBlast.Cost.Multiplier -= 175 * oldStacks
			mage.ArcaneBlast.Cost.Multiplier += 175 * newStacks

			oldMultiplier := mage.ArcaneBlastDamageMultiplier * float64(oldStacks)
			newMultiplier := mage.ArcaneBlastDamageMultiplier * float64(newStacks)
			core.Each(additiveDamageAffectedSpells, func(spell *core.Spell) {
				spell.DamageMultiplierAdditive -= oldMultiplier
				spell.DamageMultiplierAdditive += newMultiplier
			})

			if hasLivingFlameRune {
				// Living Flame is the only spell buffed multiplicatively for whatever reason
				mage.LivingFlame.DamageMultiplier /= 1 + oldMultiplier
				mage.LivingFlame.DamageMultiplier *= 1 + newMultiplier
			}
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.Flags.Matches(SpellFlagMage) && slices.Contains(affectedSpellCodes, spell.SpellCode) {
				aura.Deactivate(sim)
			}
		},
	})

	mage.ArcaneBlast = mage.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 400574},
		SpellCode:   SpellCode_MageArcaneBlast,
		SpellSchool: core.SpellSchoolArcane,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       SpellFlagMage | core.SpellFlagAPL,

		ManaCost: core.ManaCostOptions{
			BaseCost: manaCost,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: castTime,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BonusCoefficient: spellCoeff,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := sim.Roll(baseLowDamage, baseHighDamage)
			spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)

			mage.ArcaneBlastAura.Activate(sim)
			mage.ArcaneBlastAura.AddStack(sim)
		},
	})
}
