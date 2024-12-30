package human

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
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

func (s *HumanDie) Enter(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 0
	ev := event.NewHumanDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = true
}

func (s *HumanDie) Update(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse += 7

	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}
