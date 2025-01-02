package lander

import (
	"math"
	"math/rand"

	"github.com/jeffnyman/defender-redlabel/cmp"
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

func (s *LanderSearch) Enter(ai *cmp.AI, e types.IEntity) {
	sh := cmp.NewShootable()
	e.AddComponent(sh)
	dr := e.GetComponent(types.Draw).(*cmp.Draw)
	smap := dr.SpriteMap
	cl := cmp.NewCollide(smap.Frame.W/smap.Anim_frames, smap.Frame.H)
	e.AddComponent(cl)
}

func (s *LanderSearch) Update(ai *cmp.AI, e types.IEntity) {
	ai.Counter++

	pc := e.GetComponent(types.Pos).(*cmp.Pos)

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
		tc := e.GetEngine().GetPlayer().GetComponent(types.Pos).(*cmp.Pos)
		bullettime := defs.CurrentLevel().BulletTime
		dx, dy := physics.ComputeBullet(pc, tc, bullettime)
		ev := event.NewFireBullet(cmp.NewPos(pc.X, pc.Y, dx, dy))
		event.NotifyEvent(ev)
	}

	if e.Child() != e.GetID() {
		te := e.GetEngine().GetEntity(e.Child())
		tpc := te.GetComponent(types.Pos).(*cmp.Pos)

		if math.Abs(tpc.X-(pc.X+18)) < 3 {
			ai.NextState = types.LanderDrop
		}
	}
}
