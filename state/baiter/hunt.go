package baiter

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/physics"
	"github.com/jeffnyman/defender-redlabel/types"
)

type BaiterHunt struct {
	Name types.StateType
}

func NewBaiterHunt() *BaiterHunt {
	return &BaiterHunt{
		Name: types.BaiterHunt,
	}
}

func (s *BaiterHunt) GetName() types.StateType {
	return s.Name
}

func (s *BaiterHunt) Enter(ai *components.AI, e types.IEntity) {
	ai.Scratch = 0
}

func (s *BaiterHunt) Update(ai *components.AI, e types.IEntity) {
	gs := float64(defs.BaiterSpeed)
	pc := e.GetComponent(types.Pos).(*components.Pos)
	ple := e.GetEngine().GetPlayer()
	plpos := ple.GetComponent(types.Pos).(*components.Pos)

	ai.Scratch++

	if ai.Scratch == 30 {
		ai.Scratch = 0
		xoff := rand.Float64()*200 - 200
		yoff := rand.Float64()*100 - 100
		offpos := &components.Pos{X: plpos.X + xoff, Y: plpos.Y + yoff, DX: plpos.DX, DY: plpos.DY}
		pc.DX, pc.DY = physics.ComputeBullet(pc, offpos, 1)
		pc.DX = physics.Clamp(pc.DX, -gs, gs)
	}

	if !physics.OffScreen(physics.ScreenX(pc.X), pc.Y) && rand.Intn(50) == 0 {
		plp := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*components.Pos)
		bullettime := defs.CurrentLevel().BulletTime
		dx, dy := physics.ComputeBullet(pc, plp, bullettime/2)
		ev := event.NewFireBullet(components.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}
}
