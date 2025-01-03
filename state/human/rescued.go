package human

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type HumanRescued struct {
	Name types.StateType
}

func NewHumanRescued() *HumanRescued {
	return &HumanRescued{
		Name: types.HumanRescued,
	}
}

func (s *HumanRescued) GetName() types.StateType {
	return s.Name
}

func (s *HumanRescued) Enter(ai *components.AI, e types.IEntity) {
	ev := event.NewHumanRescued(e)
	event.NotifyEvent(ev)
	e.SetParent(e.GetEngine().GetPlayer().GetID())
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = 0
	pc.DY = 0
}

func (s *HumanRescued) Update(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pe := e.GetEngine().GetEntity(e.Parent())

	pec := pe.GetComponent(types.Pos).(*components.Pos)
	pc.Y = pec.Y + 50
	pc.X = pec.X

	// TODO why not aligned ?
	if pc.Y > defs.ScreenHeight-e.GetEngine().MountainHeight(pc.X) {
		ai.NextState = types.HumanWalking
		ev := event.NewHumanSaved(e)
		event.NotifyEvent(ev)
	}
}
