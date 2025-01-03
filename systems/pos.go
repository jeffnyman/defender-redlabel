package systems

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/logger"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type PosSystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewPosSystem(active bool, engine types.IEngine) *PosSystem {
	f := NewFilter()
	f.Add(types.Pos)

	return &PosSystem{
		sysname: types.PosSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (pos *PosSystem) GetName() types.SystemName {
	return pos.sysname
}

func (pos *PosSystem) Update() {
	if !pos.active {
		return
	}

	for _, e := range pos.targets {
		if e.Active() && !e.Paused() {
			pos.process(e)
		}
	}
}

func (pos *PosSystem) Draw(screen *ebiten.Image) {}

func (pos *PosSystem) process(e types.IEntity) {
	poscmp := e.GetComponent(types.Pos).(*components.Pos)

	if poscmp.Y == 9999 || poscmp.ScreenCoords {
		return
	}

	if poscmp.X < 0 {
		poscmp.X += defs.WorldWidth
	} else if poscmp.X > defs.WorldWidth {
		poscmp.X -= defs.WorldWidth
	}

	if poscmp.Y < defs.ScreenTop+20 {
		poscmp.Y = defs.ScreenTop + 20
	}

	if poscmp.Y > defs.ScreenHeight-50 {
		poscmp.Y = defs.ScreenHeight - 50
	}

	poscmp.X += poscmp.DX
	poscmp.Y += poscmp.DY
}

func (pos *PosSystem) Active() bool {
	return pos.active
}

func (pos *PosSystem) SetActive(active bool) {
	pos.active = active
}

func (pos *PosSystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := pos.targets[e.GetID()]; ok {
		return
	}

	for _, c := range pos.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}

	logger.Debug("System %T added entity %d ", pos, e.GetID())

	pos.targets[e.GetID()] = e
}

func (pos *PosSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range pos.filter.Requires() {
		if !e.HasComponent(c) {
			logger.Debug("System %T removed entity %d ", pos, e.GetID())

			delete(pos.targets, e.GetID())

			return
		}
	}
}

func (s *PosSystem) RemoveEntity(e types.IEntity) {
	logger.Debug("System %T removed entity %d ", s, e.GetID())

	delete(s.targets, e.GetID())
}
