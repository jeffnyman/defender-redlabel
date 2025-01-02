package systems

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/physics"

	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/logger"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type CollideSystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewCollideSystem(active bool, engine types.IEngine) *CollideSystem {
	f := NewFilter()
	f.Add(types.Collide)

	return &CollideSystem{
		sysname: types.CollideSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (cs *CollideSystem) GetName() types.SystemName {
	return cs.sysname
}

func (cs *CollideSystem) Update() {
	if !cs.active {
		return
	}

	pe := cs.engine.GetEntities()[defs.PlayerID]

	for _, e := range cs.targets {
		if e.Active() && !e.Paused() {
			cs.process(e, pe)
		}
	}
}

func (cs *CollideSystem) Draw(screen *ebiten.Image) {}

func (cs *CollideSystem) process(e types.IEntity, player types.IEntity) {
	ep := e.GetComponent(types.Pos).(*cmp.Pos)
	ec := e.GetComponent(types.Collide).(*cmp.Collide)
	ppos := player.GetComponent(types.Pos).(*cmp.Pos)
	psh := player.GetComponent(types.Ship).(*cmp.Ship)

	if physics.Collide(ppos.X, ppos.Y, psh.W, psh.H, ep.X, ep.Y, ec.W, ec.H) {
		ev := event.NewPlayerCollide(e)
		event.NotifyEvent(ev)
	}
}

func (cs *CollideSystem) Active() bool {
	return cs.active
}

func (cs *CollideSystem) SetActive(active bool) {
	cs.active = active
}

func (cs *CollideSystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := cs.targets[e.GetID()]; ok {
		return
	}

	for _, c := range cs.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}

	logger.Debug("System %T added entity %d ", cs, e.GetID())

	cs.targets[e.GetID()] = e
}

func (cs *CollideSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range cs.filter.Requires() {
		if !e.HasComponent(c) {
			logger.Debug("System %T removed entity %d ", cs, e.GetID())

			delete(cs.targets, e.GetID())

			return
		}
	}
}

func (cs *CollideSystem) RemoveEntity(e types.IEntity) {
	logger.Debug("System %T removed entity %d ", cs, e.GetID())

	delete(cs.targets, e.GetID())
}
