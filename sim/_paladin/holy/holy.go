package holy

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/paladin"
)

func RegisterHolyPaladin() {
	core.RegisterAgentFactory(
		proto.Player_HolyPaladin{},
		proto.Spec_SpecHolyPaladin,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewHolyPaladin(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_HolyPaladin) // I don't really understand this line
			if !ok {
				panic("Invalid spec value for Holy Paladin!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewHolyPaladin(character *core.Character, options *proto.Player) *HolyPaladin {
	holyOptions := options.GetHolyPaladin()

	holy := &HolyPaladin{
		Paladin: paladin.NewPaladin(character, options, holyOptions.Options.Aura),
		Options: holyOptions.Options,
	}

	holy.PaladinAura = holyOptions.Options.Aura

	return holy
}

type HolyPaladin struct {
	*paladin.Paladin

	Options *proto.HolyPaladin_Options
}

func (holy *HolyPaladin) GetPaladin() *paladin.Paladin {
	return holy.Paladin
}

func (holy *HolyPaladin) Initialize() {
	holy.Paladin.Initialize()
}

func (holy *HolyPaladin) Reset(sim *core.Simulation) {
	holy.Paladin.Reset(sim)
}
