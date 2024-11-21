import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Consumes,
	Debuffs,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SaygesFortune,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import { TankWarrior_Options as TankWarriorOptions, WarriorShout, WarriorStance } from '../core/proto/warrior.js';
import Phase4APL from './apls/phase_4.apl.json';
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

export const APLPhase4 = PresetUtils.makePresetAPLRotation('P4 Prot/Fury', Phase4APL, { customCondition: player => player.getLevel() === 60 });

export const APLPresets = {
	[Phase.Phase1]: [],
	[Phase.Phase2]: [],
	[Phase.Phase3]: [],
	[Phase.Phase4]: [APLPhase4],
	[Phase.Phase5]: [],
};

export const DefaultAPLs: Record<number, PresetUtils.PresetRotation> = {
	25: APLPresets[Phase.Phase1][0],
	40: APLPresets[Phase.Phase2][0],
	50: APLPresets[Phase.Phase3][0],
	60: APLPresets[Phase.Phase4][0],
};

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsPhase4Prot = PresetUtils.makePresetTalents('60 Prot', SavedTalents.create({ talentsString: '20304300302-03-55200110530201051' }), {
	customCondition: player => player.getLevel() === 60,
});

export const TalentsPhase4Fury = PresetUtils.makePresetTalents('60 Fury', SavedTalents.create({ talentsString: '33302300302-05050005505010051' }), {
	customCondition: player => player.getLevel() === 60,
});

export const TalentPresets = {
	[Phase.Phase1]: [],
	[Phase.Phase2]: [],
	[Phase.Phase3]: [],
	[Phase.Phase4]: [TalentsPhase4Prot, TalentsPhase4Fury],
	[Phase.Phase5]: [],
};

export const DefaultTalents = TalentPresets[Phase.Phase4][0];

export const PresetBuildTanky = PresetUtils.makePresetBuild('Tanky', { gear: DefaultGear, talents: TalentsPhase4Prot, rotation: DefaultAPLs[60] });
export const PresetBuildDamage = PresetUtils.makePresetBuild('Damage', { gear: DefaultGear, talents: TalentsPhase4Fury, rotation: DefaultAPLs[60] });

///////////////////////////////////////////////////////////////////////////
//                                 Options Presets
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = TankWarriorOptions.create({
	startingRage: 0,
	shout: WarriorShout.WarriorShoutBattle,
	stance: WarriorStance.WarriorStanceDefensive,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultPotion: Potions.MightyRagePotion,
	dragonBreathChili: true,
	food: Food.FoodSmokedDesertDumpling,
	flask: Flask.FlaskOfTheTitans,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.Windfury,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	battleShout: TristateEffect.TristateEffectImproved,
	devotionAura: TristateEffect.TristateEffectRegular,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	stoneskinTotem: TristateEffect.TristateEffectRegular,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	fengusFerocity: true,
	moldarsMoxie: true,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	warchiefsBlessing: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	giftOfArthas: true,
	improvedScorch: true,
	insectSwarm: true,
	sunderArmor: true,
});

export const OtherDefaults = {
	profession1: Profession.Blacksmithing,
	profession2: Profession.Enchanting,
	race: Race.RaceHuman,
};
