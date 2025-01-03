package lander

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type LanderDie struct {
	Name types.StateType
}

func NewLanderDie() *LanderDie {
	return &LanderDie{
		Name: types.LanderDie,
	}
}

func (s *LanderDie) GetName() types.StateType {
	return s.Name
}

func (s *LanderDie) Enter(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse = 0
	ev := event.NewLanderDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*components.RadarDraw)
	rdc.Hide = true
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = 0
	pc.DY = 0
	e.RemoveComponent(types.Collide)
}

func (s *LanderDie) Update(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse += 7

	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}
