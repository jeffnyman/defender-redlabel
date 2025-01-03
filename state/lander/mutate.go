package lander

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/graphics"
	"github.com/jeffnyman/defender-redlabel/physics"
	"github.com/jeffnyman/defender-redlabel/types"
)

type LanderMutate struct {
	Name types.StateType
}

func NewLanderMutate() *LanderMutate {
	return &LanderMutate{
		Name: types.LanderMutate,
	}
}

func (s *LanderMutate) GetName() types.StateType {
	return s.Name
}

func (s *LanderMutate) Enter(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	pc.DY = 0
	pc.DX = 0
	dc := e.GetComponent(types.Draw).(*components.Draw)
	dc.SpriteMap = graphics.GetSpriteMap("mutant.png")
	rc := e.GetComponent(types.RadarDraw).(*components.RadarDraw)
	rc.Cycle = true
	ai.Counter = 0
	ai.Scratch = 0
}

func (s *LanderMutate) Update(ai *components.AI, e types.IEntity) {
	gs := float64(defs.LanderSpeed)
	pc := e.GetComponent(types.Pos).(*components.Pos)
	ppc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*components.Pos)

	if pc.X > ppc.X {
		pc.DX = -gs * 3
	} else {
		pc.DX = gs * 3
	}

	if pc.Y > ppc.Y {
		pc.DY = -gs * 2
	} else {
		pc.DY = gs * 2
	}

	ai.Counter++

	if ai.Counter > 2 {
		ai.Counter = 0

		pc.DY = physics.RandChoiceF([]float64{-gs, 0, gs})
		pc.X += physics.RandChoiceF([]float64{-20, 0, 20})
		pc.Y += physics.RandChoiceF([]float64{-20, 0, 20})
	}

	ai.Scratch++

	if ai.Scratch > 7 {
		ai.Scratch = 0
		ev := event.NewMutantSound(e)
		event.NotifyEvent(ev)
	}

	// TODO defs bullet rate
	if !physics.OffScreen(physics.ScreenX(pc.X), pc.Y) && rand.Intn(100) == 0 {
		tc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*components.Pos)
		bullettime := defs.CurrentLevel().BulletTime
		dx, dy := physics.ComputeBullet(pc, tc, bullettime)
		ev := event.NewFireBullet(components.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

}
