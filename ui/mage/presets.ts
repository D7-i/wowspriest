import { Phase } from '../core/constants/other';
import * as PresetUtils from '../core/preset_utils';
import {
	Conjured,
	Consumes,
	Debuffs,
	EnchantedSigil,
	FirePowerBuff,
	Flask,
	Food,
	FrostPowerBuff,
	IndividualBuffs,
	MageScroll,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common';
import { Mage_Options as MageOptions, Mage_Options_ArmorType as ArmorType } from '../core/proto/mage';
import { SavedTalents } from '../core/proto/ui';
import Phase1APLArcane from './apls/p1_arcane.apl.json';
import Phase1APLFire from './apls/p1_fire.apl.json';
import Phase2APLArcane from './apls/p2_arcane.apl.json';
import Phase2APLFire from './apls/p2_fire.apl.json';
// import Phase3APLArcane from './apls/p3_arcane.apl.json';
import Phase3APLFire from './apls/p3_fire.apl.json';
import Phase3APLFrost from './apls/p3_frost.apl.json';
import Phase4APLArcane from './apls/p4_arcane.apl.json';
import Phase4APLFire from './apls/p4_fire.apl.json';
import Phase4APLFrost from './apls/p4_frost.apl.json';
import Phase5APLFire from './apls/p5_fire.apl.json';
import Phase5APLSpellFrost from './apls/p5_spellfrost.apl.json';
import BlankGear from './gear_sets/blank.gear.json';

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);

export const GearPresets = {};

export const DefaultGear = GearBlank;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLArcanePhase1 = PresetUtils.makePresetAPLRotation('P1 Arcane', Phase1APLArcane, {
	customCondition: player => player.getLevel() === 25,
});
export const APLFirePhase1 = PresetUtils.makePresetAPLRotation('P1 Fire', Phase1APLFire, {
	customCondition: player => player.getLevel() === 25,
});

export const APLArcanePhase2 = PresetUtils.makePresetAPLRotation('P2 Arcane', Phase2APLArcane, {
	customCondition: player => player.getLevel() === 40,
});
export const APLFirePhase2 = PresetUtils.makePresetAPLRotation('P2 Fire', Phase2APLFire, {
	customCondition: player => player.getLevel() === 40,
});

// No new Phase 3 Arcane presets at the moment
export const APLArcanePhase3 = APLArcanePhase2;
export const APLFirePhase3 = PresetUtils.makePresetAPLRotation('P3 Fire', Phase3APLFire, {
	customCondition: player => player.getLevel() === 50,
});
export const APLFrostPhase3 = PresetUtils.makePresetAPLRotation('P3 Frost', Phase3APLFrost, {
	customCondition: player => player.getLevel() === 50,
});

export const APLArcanePhase4 = PresetUtils.makePresetAPLRotation('P4 Arcane', Phase4APLArcane, {
	customCondition: player => player.getLevel() >= 60,
});
export const APLFirePhase4 = PresetUtils.makePresetAPLRotation('P4 Fire', Phase4APLFire, {
	customCondition: player => player.getLevel() >= 60,
});
export const APLFrostPhase4 = PresetUtils.makePresetAPLRotation('P4 Frost', Phase4APLFrost, {
	customCondition: player => player.getLevel() >= 60,
});

export const APLFirePhase5 = PresetUtils.makePresetAPLRotation('P5 Fire', Phase5APLFire, {
	customCondition: player => player.getLevel() >= 60,
});
export const APLSpellfrostPhase5 = PresetUtils.makePresetAPLRotation('P5 Frost', Phase5APLSpellFrost, {
	customCondition: player => player.getLevel() >= 60,
});

export const APLPresets = {
	[Phase.Phase1]: [APLArcanePhase1, APLFirePhase1, APLFirePhase1],
	[Phase.Phase2]: [APLArcanePhase2, APLFirePhase2, APLFirePhase2],
	[Phase.Phase3]: [APLArcanePhase3, APLFirePhase3, APLFrostPhase3],
	[Phase.Phase4]: [APLArcanePhase4, APLFirePhase4, APLFrostPhase4],
	[Phase.Phase5]: [APLFirePhase5, APLSpellfrostPhase5],
};

export const DefaultAPLs: Record<number, Record<number, PresetUtils.PresetRotation>> = {
	25: {
		0: APLPresets[Phase.Phase1][0],
		1: APLPresets[Phase.Phase1][1],
		2: APLPresets[Phase.Phase1][2],
	},
	40: {
		0: APLPresets[Phase.Phase2][0],
		1: APLPresets[Phase.Phase2][1],
		// Normally frost but frost is unfortunately just too bad to warrant including for now
		2: APLPresets[Phase.Phase2][2],
		// Frostfire
		3: APLPresets[Phase.Phase2][2],
	},
	50: {
		0: APLPresets[Phase.Phase3][0],
		1: APLPresets[Phase.Phase3][1],
		2: APLPresets[Phase.Phase3][2],
	},
	60: {
		0: APLPresets[Phase.Phase5][1],
		1: APLPresets[Phase.Phase5][0],
		2: APLPresets[Phase.Phase5][1],
	},
};

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsArcanePhase1 = PresetUtils.makePresetTalents('25 Arcane', SavedTalents.create({ talentsString: '22500502' }), {
	customCondition: player => player.getLevel() === 25,
});
export const TalentsFirePhase1 = PresetUtils.makePresetTalents('25 Fire', SavedTalents.create({ talentsString: '-5050020121' }), {
	customCondition: player => player.getLevel() === 25,
});

export const TalentsArcanePhase2 = PresetUtils.makePresetTalents('40 Arcane', SavedTalents.create({ talentsString: '2250050310031531' }), {
	customCondition: player => player.getLevel() === 40,
});
export const TalentsFirePhase2 = PresetUtils.makePresetTalents('40 Fire', SavedTalents.create({ talentsString: '-5050020123033151' }), {
	customCondition: player => player.getLevel() === 40,
});

// No new Phase 3 Arcane presets at the moment
export const TalentsArcanePhase3 = TalentsArcanePhase2;
export const TalentsFirePhase3 = PresetUtils.makePresetTalents('50 Fire', SavedTalents.create({ talentsString: '-0550020123033151-2035' }), {
	customCondition: player => player.getLevel() === 50,
});
export const TalentsFrostPhase3 = PresetUtils.makePresetTalents('50 Frost', SavedTalents.create({ talentsString: '-055-20350203100351051' }), {
	customCondition: player => player.getLevel() === 50,
});

export const TalentsArcanePhase4 = PresetUtils.makePresetTalents('60 Arcane', SavedTalents.create({ talentsString: '0550050210031531-054-203500001' }), {
	customCondition: player => player.getLevel() === 60,
});
export const TalentsFirePhase4 = PresetUtils.makePresetTalents('60 Fire', SavedTalents.create({ talentsString: '21-5052300123033151-203500031' }), {
	customCondition: player => player.getLevel() === 60,
});
export const TalentsFrostfirePhase4 = PresetUtils.makePresetTalents('60 Frostfire', SavedTalents.create({ talentsString: '-0550320003021-2035020310035105' }), {
	customCondition: player => player.getLevel() === 60,
});

export const TalentsArcanePhase5 = PresetUtils.makePresetTalents('60 Arcane', SavedTalents.create({ talentsString: '2500550010031531--2035020310004' }), {
	customCondition: player => player.getLevel() === 60,
});
export const TalentsFrostPhase5 = PresetUtils.makePresetTalents('60 Frost', SavedTalents.create({ talentsString: '250025001002--05350203100351051' }), {
	customCondition: player => player.getLevel() === 60,
});

export const TalentPresets = {
	[Phase.Phase1]: [TalentsArcanePhase1, TalentsFirePhase1, TalentsFirePhase1],
	[Phase.Phase2]: [TalentsArcanePhase2, TalentsFirePhase2, TalentsFirePhase2],
	[Phase.Phase3]: [TalentsArcanePhase3, TalentsFirePhase3, TalentsFrostPhase3],
	[Phase.Phase4]: [TalentsFrostfirePhase4],
	[Phase.Phase5]: [TalentsArcanePhase5, TalentsFirePhase4, TalentsFrostPhase5],
};

export const DefaultTalentsArcane = TalentPresets[Phase.Phase5][0];
export const DefaultTalentsFire = TalentPresets[Phase.Phase5][1];
export const DefaultTalentsFrostfire = TalentPresets[Phase.Phase4][0];
export const DefaultTalentsFrost = TalentPresets[Phase.Phase5][2];

export const DefaultTalents = DefaultTalentsFire;

export const PresetBuildArcane = PresetUtils.makePresetBuild('Arcane', {
	gear: DefaultGear,
	talents: DefaultTalentsArcane,
	rotation: DefaultAPLs[60][0],
});
export const PresetBuildFire = PresetUtils.makePresetBuild('Fire', { gear: DefaultGear, talents: DefaultTalentsFire, rotation: DefaultAPLs[60][1] });
export const PresetBuildFrost = PresetUtils.makePresetBuild('Frost', { gear: DefaultGear, talents: DefaultTalentsFrost, rotation: DefaultAPLs[60][2] });

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = MageOptions.create({
	armor: ArmorType.MoltenArmor,
});

export const DefaultConsumes = Consumes.create({
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	enchantedSigil: EnchantedSigil.FlowingWatersSigil,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	frostPowerBuff: FrostPowerBuff.ElixirOfFrostPower,
	mageScroll: MageScroll.MageScrollArcanePower,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	mildlyIrradiatedRejuvPot: true,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	aspectOfTheLion: true,
	demonicPact: 110,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
	vampiricTouch: 300,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	mightOfStormwind: true,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: true,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	valorOfAzeroth: true,
	warchiefsBlessing: true,
});

export const DefaultDebuffs = Debuffs.create({
	dreamstate: true,
	improvedFaerieFire: true,
	improvedScorch: true,
	judgementOfWisdom: true,
	markOfChaos: true,
	occultPoison: true,
	wintersChill: true,
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Alchemy,
	profession2: Profession.Tailoring,
};
