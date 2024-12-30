package pod

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type PodDie struct {
	Name types.StateType
}

func NewPodDie() *PodDie {
	return &PodDie{
		Name: types.PodDie,
	}
}

func (s *PodDie) GetName() types.StateType {
	return s.Name
}

func (s *PodDie) Enter(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse = 0
	ev := event.NewPodDie(e)
	event.NotifyEvent(ev)
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = true
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
	e.RemoveComponent(types.Collide)
}

func (s *PodDie) Update(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse += 7

	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}