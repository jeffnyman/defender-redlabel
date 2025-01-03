package lander

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/types"
)

type LanderWait struct {
	Name types.StateType
}

func NewLanderWait() *LanderWait {
	return &LanderWait{
		Name: types.LanderWait,
	}
}

func (s *LanderWait) GetName() types.StateType {
	return s.Name
}

func (s *LanderWait) Enter(ai *components.AI, e types.IEntity) {
	dr := e.GetComponent(types.Draw).(*components.Draw)
	dr.Hide = true
	rdc := e.GetComponent(types.RadarDraw).(*components.RadarDraw)
	rdc.Hide = true
}

func (s *LanderWait) Update(ai *components.AI, e types.IEntity) {
	ai.Wait--

	if ai.Wait <= 0 {
		ai.NextState = types.LanderMaterialize
	}
}
