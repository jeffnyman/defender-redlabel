package lander

import (
	"math"

	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type LanderDrop struct {
	Name types.StateType
}

func NewLanderDrop() *LanderDrop {
	return &LanderDrop{
		Name: types.LanderDrop,
	}
}

func (s *LanderDrop) GetName() types.StateType {
	return s.Name
}

func (s *LanderDrop) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = 0
	pc.DY = 1.2 * defs.LanderSpeed
	ai.Counter = 0
	te := e.GetEngine().GetEntity(e.Child())
	tpc := te.GetComponent(types.Pos).(*components.Pos)
	tpc.DX = 0
}

func (s *LanderDrop) Update(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	te := e.GetEngine().GetEntity(e.Child())
	tpc := te.GetComponent(types.Pos).(*components.Pos)

	if math.Abs(pc.Y-tpc.Y) < 5 {
		ai.NextState = types.LanderGrab
		tai := te.GetComponent(types.AI).(*components.AI)
		tai.NextState = types.HumanGrabbed
	}
}
