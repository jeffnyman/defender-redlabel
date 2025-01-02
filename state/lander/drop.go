package lander

import (
	"math"

	"github.com/jeffnyman/defender-redlabel/cmp"
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

func (s *LanderDrop) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 1.2 * defs.LanderSpeed
	ai.Counter = 0
	te := e.GetEngine().GetEntity(e.Child())
	tpc := te.GetComponent(types.Pos).(*cmp.Pos)
	tpc.DX = 0
}

func (s *LanderDrop) Update(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	te := e.GetEngine().GetEntity(e.Child())
	tpc := te.GetComponent(types.Pos).(*cmp.Pos)

	if math.Abs(pc.Y-tpc.Y) < 5 {
		ai.NextState = types.LanderGrab
		tai := te.GetComponent(types.AI).(*cmp.AI)
		tai.NextState = types.HumanGrabbed
	}
}
