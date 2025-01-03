package player

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/graphics"
	"github.com/jeffnyman/defender-redlabel/types"
)

type PlayerDie struct {
	Name types.StateType
}

func NewPlayerDie() *PlayerDie {
	return &PlayerDie{
		Name: types.PlayerDie,
	}
}

func (s *PlayerDie) GetName() types.StateType {
	return s.Name
}

func (s *PlayerDie) Enter(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.SpriteMap = graphics.GetSpriteMap("shipd.png")
	dc.Frame = 0
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = 0
	pc.DY = 0
	ai.Counter = 0
	fle := e.GetEngine().GetEntity(e.Child())
	fdc := fle.GetComponent(types.Draw).(*components.Draw)
	fdc.Hide = true

}

func (s *PlayerDie) Update(ai *components.AI, e types.IEntity) {
	ai.Counter++
	dc := e.GetComponent(types.Draw).(*components.Draw)

	if ai.Counter == 60 {
		dc.Hide = true
		pc := e.GetComponent(types.Pos).(*components.Pos)
		e.GetEngine().TriggerPS(pc.X, pc.Y)
		ev := event.NewPlayerExplode(e)
		event.NotifyEvent(ev)
	}

	if ai.Counter == 180 {
		defs.PlayerLives--
		ai.NextState = types.PlayerPlay
		dc.Hide = false
		dc.SpriteMap = graphics.GetSpriteMap("ship.png")
		dc.Frame = 0
	}
}
