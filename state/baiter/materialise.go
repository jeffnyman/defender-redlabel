package baiter

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BaiterMaterialize struct {
	Name types.StateType
}

func NewBaiterMaterialize() *BaiterMaterialize {
	return &BaiterMaterialize{
		Name: types.BaiterMaterialize,
	}
}

func (s *BaiterMaterialize) GetName() types.StateType {
	return s.Name
}

func (s *BaiterMaterialize) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DY = 0
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Hide = false
	dc.Disperse = 300
	rdc := e.GetComponent(types.RadarDraw).(*components.RadarDraw)
	rdc.Hide = false
	ev := event.NewMaterialize(e)
	event.NotifyEvent(ev)
}

func (s *BaiterMaterialize) Update(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = e.GetEngine().GetPlayer().GetComponent(types.Pos).(*components.Pos).DX
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse -= 5

	if dc.Disperse < 10 {
		dc.Disperse = 0
		ai.NextState = types.BaiterHunt
	}
}
