package bomber

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/gl"
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

func (s *BomberMove) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DX = gl.BomberSpeed
	pc.DY = -gl.BomberSpeed
}

func (s *BomberMove) Update(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)

	if pc.Y < gl.ScreenTop+50 || pc.Y > gl.ScreenHeight-100 {
		pc.DY = -pc.DY
	}

	if rand.Intn(40) == 0 {
		e.GetEngine().TriggerBomb(pc.X, pc.Y)
	}
}
