package lander

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type LanderMaterialize struct {
	Name types.StateType
}

func NewLanderMaterialize() *LanderMaterialize {
	return &LanderMaterialize{
		Name: types.LanderMaterialize,
	}
}

func (s *LanderMaterialize) GetName() types.StateType {
	return s.Name
}

func (s *LanderMaterialize) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	pc.DY = 0
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Hide = false
	dc.Disperse = 300
	rdc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rdc.Hide = false
	ev := event.NewMaterialize(e)
	event.NotifyEvent(ev)

	ai.Counter = 0

	for _, id := range e.GetEngine().GetActiveEntitiesOfClass(types.Human) {
		hum := e.GetEngine().GetEntity(id)

		if hum.Parent() == hum.GetID() {
			humpos := hum.GetComponent(types.Pos).(*cmp.Pos)

			if humpos.X-pc.X < 4000 {
				e.SetChild(hum.GetID())
				hum.SetParent(e.GetID())
				pc.X = humpos.X - rand.Float64()*2*defs.ScreenWidth
				break
			}
		}
	}
}

func (s *LanderMaterialize) Update(ai *cmp.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.Disperse -= 5

	if dc.Disperse < 10 {
		dc.Disperse = 0
		ai.NextState = types.LanderSearch
	}
}
