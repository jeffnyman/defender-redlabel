package baiter

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BaiterDie struct {
	Name types.StateType
}

func NewBaiterDie() *BaiterDie {
	return &BaiterDie{
		Name: types.BaiterDie,
	}
}

func (s *BaiterDie) GetName() types.StateType {
	return s.Name
}

func (s *BaiterDie) Enter(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 0
	ev := event.NewBaiterDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = true
	e.RemoveComponent(types.Collide)
}

func (s *BaiterDie) Update(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse += 7

	if dc.Disperse > 300 {
		ai.NextState = types.BaiterWait
	}
}
