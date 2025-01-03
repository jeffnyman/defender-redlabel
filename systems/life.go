package systems

import (
	"github.com/jeffnyman/defender-redlabel/components"

	"github.com/jeffnyman/defender-redlabel/logger"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type LifeSystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewLifeSystem(active bool, engine types.IEngine) *LifeSystem {
	f := NewFilter()
	f.Add(types.Life)

	return &LifeSystem{
		sysname: types.LifeSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (pos *LifeSystem) GetName() types.SystemName {
	return pos.sysname
}

func (pos *LifeSystem) Update() {
	if !pos.active {
		return
	}

	for _, e := range pos.targets {
		if e.Active() && !e.Paused() {
			pos.process(e)
		}
	}
}

func (pos *LifeSystem) Draw(screen *ebiten.Image) {}

func (pos *LifeSystem) process(e types.IEntity) {
	cmp := e.GetComponent(types.Life).(*components.Life)
	cmp.TicksToLive--

	if cmp.TicksToLive < 0 {
		e.SetActive(false)
	}
}

func (pos *LifeSystem) Active() bool {
	return pos.active
}

func (pos *LifeSystem) SetActive(active bool) {
	pos.active = active
}

func (pos *LifeSystem) AddEntityIfRequired(e types.IEntity) {
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

func (pos *LifeSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range pos.filter.Requires() {
		if !e.HasComponent(c) {
			logger.Debug("System %T removed entity %d ", pos, e.GetID())

			delete(pos.targets, e.GetID())

			return
		}
	}
}

func (s *LifeSystem) RemoveEntity(e types.IEntity) {
	logger.Debug("System %T removed entity %d ", s, e.GetID())

	delete(s.targets, e.GetID())
}
