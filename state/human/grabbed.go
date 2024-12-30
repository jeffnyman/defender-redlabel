package human

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
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

func (s *HumanGrabbed) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = 0
	ev := event.NewHumanGrabbed(e)
	event.NotifyEvent(ev)
}

func (s *HumanGrabbed) Update(ai *cmp.AI, e types.IEntity) {
	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pe := e.GetEngine().GetEntity(e.Parent())
	pai := pe.GetComponent(types.AI).(*cmp.AI)

	if pai.State != types.LanderDie {
		pec := pe.GetComponent(types.Pos).(*cmp.Pos)
		pc.Y = pec.Y + 50
	} else {
		ai.NextState = types.HumanDropping
	}
}
