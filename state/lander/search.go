package lander

import (
	"math"
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/physics"
	"github.com/jeffnyman/defender-redlabel/types"
)

type LanderSearch struct {
	Name types.StateType
}

func NewLanderSearch() *LanderSearch {
	return &LanderSearch{
		Name: types.LanderSearch,
	}
}

func (s *LanderSearch) GetName() types.StateType {
	return s.Name
}

func (s *LanderSearch) Enter(ai *components.AI, e types.IEntity) {
	sh := components.NewShootable()
	e.AddComponent(sh)
	dr := e.GetComponent(types.Draw).(*components.Draw)
	smap := dr.SpriteMap
	cl := components.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	e.AddComponent(cl)
}

func (s *LanderSearch) Update(ai *components.AI, e types.IEntity) {
	ai.Counter++

	pc := e.GetComponent(types.Pos).(*components.Pos)

	if ai.Counter > 5 {
		ai.Counter = 0
		mh := e.GetEngine().MountainHeight(pc.X)

		if pc.Y+200 < defs.ScreenHeight-mh {
			ai.Scratch++
		} else {
			ai.Scratch--
		}
	}

	if ai.Scratch < 0 {
		ai.Scratch = 0
	}

	if ai.Scratch > 5 {
		ai.Scratch = 5
	}

	switch ai.Scratch {
	case 0:
		pc.DY = -defs.LanderSpeed
	case 1, 2, 3, 4:
		pc.DY = 0
	case 5:
		pc.DY = defs.LanderSpeed
	}

	// TODO defs bullet rate
	if !physics.OffScreen(physics.ScreenX(pc.X), pc.Y) && rand.Intn(100) == 0 {
		tc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*components.Pos)
		bullettime := defs.CurrentLevel().BulletTime
		dx, dy := physics.ComputeBullet(pc, tc, bullettime)
		ev := event.NewFireBullet(components.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

	if e.Child() != e.GetID() {
		te := e.GetEngine().GetEntity(e.Child())
		tpc := te.GetComponent(types.Pos).(*components.Pos)

		if math.Abs(tpc.X-(pc.X+18)) < 3 {
			ai.NextState = types.LanderDrop
		}
	}
}
