package baiter

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BaiterWait struct {
	Name types.StateType
}

func NewBaiterWait() *BaiterWait {
	return &BaiterWait{
		Name: types.BaiterWait,
	}
}

func (s *BaiterWait) GetName() types.StateType {
	return s.Name
}

func (s *BaiterWait) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.Y = 9999
}

func (s *BaiterWait) Update(ai *cmp.AI, e types.IEntity) {
	if defs.CurrentLevel().LanderCount-defs.LandersKilled < 3 {
		ai.NextState = types.BaiterMaterialize
		pc := e.GetComponent(types.Pos).(*cmp.Pos)
		pc.Y = defs.ScreenHeight / 2
		pc.X = defs.CameraX() + rand.Float64()*defs.ScreenWidth
		pc.DX = e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos).DX
	}
}
