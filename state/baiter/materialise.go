package baiter

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BaiterMaterialise struct {
	Name types.StateType
}

func NewBaiterMaterialise() *BaiterMaterialise {
	return &BaiterMaterialise{
		Name: types.BaiterMaterialise,
	}
}

func (s *BaiterMaterialise) GetName() types.StateType {
	return s.Name
}

func (s *BaiterMaterialise) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DY = 0
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Hide = false
	dc.Disperse = 300
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = false
	ev := event.NewMaterialise(e)
	event.NotifyEvent(ev)
}

func (s *BaiterMaterialise) Update(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos).DX
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse -= 5

	if dc.Disperse < 10 {
		dc.Disperse = 0
		ai.NextState = types.BaiterHunt
	}
}
