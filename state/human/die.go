package human

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type HumanDie struct {
	Name types.StateType
}

func NewHumanDie() *HumanDie {
	return &HumanDie{
		Name: types.HumanDie,
	}
}

func (s *HumanDie) GetName() types.StateType {
	return s.Name
}

func (s *HumanDie) Enter(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse = 0
	ev := event.NewHumanDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*components.RadarDraw)
	rdc.Hide = true
}

func (s *HumanDie) Update(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse += 7

	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}
