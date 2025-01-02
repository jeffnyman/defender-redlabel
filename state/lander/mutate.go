package lander

import (
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/gl"
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

func (s *LanderMutate) Enter(ai *cmp.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	pc.DY = 0
	pc.DX = 0
	dc := e.GetComponent(types.Draw).(*cmp.Draw)
	dc.SpriteMap = graphics.GetSpriteMap("mutant.png")
	rc := e.GetComponent(types.RadarDraw).(*cmp.RadarDraw)
	rc.Cycle = true
	ai.Counter = 0
	ai.Scratch = 0
}

func (s *LanderMutate) Update(ai *cmp.AI, e types.IEntity) {
	gs := float64(gl.LanderSpeed)
	pc := e.GetComponent(types.Pos).(*cmp.Pos)
	ppc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos)

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

	// TODO gl bullet rate
	if !physics.OffScreen(physics.ScreenX(pc.X), pc.Y) && rand.Intn(100) == 0 {
		tc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos)
		bullettime := gl.CurrentLevel().BulletTime
		dx, dy := physics.ComputeBullet(pc, tc, bullettime)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

}
