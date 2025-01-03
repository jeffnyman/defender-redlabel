package bomber

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BomberMove struct {
	Name types.StateType
}

func NewBomberMove() *BomberMove {
	return &BomberMove{
		Name: types.BomberMove,
	}
}

func (s *BomberMove) GetName() types.StateType {
	return s.Name
}

func (s *BomberMove) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DX = defs.BomberSpeed
	pc.DY = -defs.BomberSpeed
}

func (s *BomberMove) Update(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)

	if pc.Y < defs.ScreenTop+50 || pc.Y > defs.ScreenHeight-100 {
		pc.DY = -pc.DY
	}

	if rand.Intn(40) == 0 {
		e.GetEngine().TriggerBomb(pc.X, pc.Y)
	}
}
