package human

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type HumanGrabbed struct {
	Name types.StateType
}

func NewHumanGrabbed() *HumanGrabbed {
	return &HumanGrabbed{
		Name: types.HumanGrabbed,
	}
}

func (s *HumanGrabbed) GetName() types.StateType {
	return s.Name
}

func (s *HumanGrabbed) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = 0
	ev := event.NewHumanGrabbed(e)
	event.NotifyEvent(ev)
}

func (s *HumanGrabbed) Update(ai *components.AI, e types.IEntity) {
	ai.Counter++

	pc := e.GetComponent(types.Pos).(*components.Pos)
	pe := e.GetEngine().GetEntity(e.Parent())
	pai := pe.GetComponent(types.AI).(*components.AI)

	if pai.State != types.LanderDie {
		pec := pe.GetComponent(types.Pos).(*components.Pos)
		pc.Y = pec.Y + 50
	} else {
		ai.NextState = types.HumanDropping
	}
}
