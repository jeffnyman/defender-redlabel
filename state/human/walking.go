package human

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type HumanWalking struct {
	Name types.StateType
}

func NewHumanWalking() *HumanWalking {
	return &HumanWalking{
		Name: types.HumanWalking,
	}
}

func (s *HumanWalking) GetName() types.StateType {
	return s.Name
}

func (s *HumanWalking) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = defs.HumanSpeed
	pc.DY = 0
	ai.Counter = 0
	pc.Y = defs.ScreenHeight - e.GetEngine().MountainHeight(pc.X)
}

func (s *HumanWalking) Update(ai *components.AI, e types.IEntity) {
	ai.Counter++

	pc := e.GetComponent(types.Pos).(*components.Pos)

	pc.Y = defs.ScreenHeight - e.GetEngine().MountainHeight(pc.X)
}
