package lander

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/components"
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

func (s *LanderMaterialize) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = 0
	pc.DY = 0
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Hide = false
	dc.Disperse = 300
	rdc := e.GetComponent(types.RadarDraw).(*components.RadarDraw)
	rdc.Hide = false
	ev := event.NewMaterialize(e)
	event.NotifyEvent(ev)

	ai.Counter = 0

	for _, id := range e.GetEngine().GetActiveEntitiesOfClass(types.Human) {
		hum := e.GetEngine().GetEntity(id)

		if hum.Parent() == hum.GetID() {
			humpos := hum.GetComponent(types.Pos).(*components.Pos)

			if humpos.X-pc.X < 4000 {
				e.SetChild(hum.GetID())
				hum.SetParent(e.GetID())
				pc.X = humpos.X - rand.Float64()*2*defs.ScreenWidth
				break
			}
		}
	}
}

func (s *LanderMaterialize) Update(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse -= 5

	if dc.Disperse < 10 {
		dc.Disperse = 0
		ai.NextState = types.LanderSearch
	}
}
