package bomber

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BomberDie struct {
	Name types.StateType
}

func NewBomberDie() *BomberDie {
	return &BomberDie{
		Name: types.BomberDie,
	}
}

func (s *BomberDie) GetName() types.StateType {
	return s.Name
}

func (s *BomberDie) Enter(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse = 0
	ev := event.NewBomberDie(e)
	event.NotifyEvent(ev)
	e.RemoveComponent(types.Collide)
}

func (s *BomberDie) Update(ai *components.AI, e types.IEntity) {
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.Disperse += 7

	if dc.Disperse > 300 {
		e.SetActive(false)
	}
}
