package systems

import (
	"github.com/jeffnyman/defender-redlabel/components"

	"math"

	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/logger"
	"github.com/jeffnyman/defender-redlabel/physics"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type LaserMoveSystem struct {
	sysname types.SystemName
	filter  *Filter
	active  bool
	engine  types.IEngine
	targets map[types.EntityID]types.IEntity
}

func NewLaserMoveSystem(active bool, engine types.IEngine) *LaserMoveSystem {
	f := NewFilter()
	f.Add(types.LaserMove)
	f.Add(types.Pos)

	return &LaserMoveSystem{
		sysname: types.LaserMoveSystem,
		active:  active,
		filter:  f,
		engine:  engine,
		targets: make(map[types.EntityID]types.IEntity),
	}
}

func (lms *LaserMoveSystem) GetName() types.SystemName {
	return lms.sysname
}

func (lms *LaserMoveSystem) Update() {
	if !lms.active {
		return
	}

	for _, e := range lms.targets {
		if e.Active() {
			lms.process(e)
		}
	}
}

func (lms *LaserMoveSystem) process(laserEnt types.IEntity) {
	pe := laserEnt.GetEngine().GetEntity(defs.PlayerID)
	ppc := pe.GetComponent(types.Pos).(*components.Pos)
	psc := pe.GetComponent(types.Ship).(*components.Ship)
	lpc := laserEnt.GetComponent(types.Pos).(*components.Pos)
	lmc := laserEnt.GetComponent(types.LaserMove).(*components.LaserMove)
	lpc.X += lpc.DX * (20 + math.Abs(ppc.DX))
	lmc.Length += 70

	var h2 float64 = 4
	y2 := lpc.Y
	x2 := physics.ScreenX(lpc.X)
	w2 := lmc.Length

	if psc.Direction == -1 {
		x2 = x2 - lmc.Length
	}

	for _, v := range laserEnt.GetEngine().GetEntitiesWithComponent(types.Shootable) {
		tpc := v.GetComponent(types.Pos).(*components.Pos)
		x1 := physics.ScreenX(tpc.X)
		y1 := tpc.Y

		if physics.OffScreen(x1, tpc.Y) {
			continue
		}

		if v.HasComponent(types.Collide) {
			tcc := v.GetComponent(types.Collide).(*components.Collide)
			w1 := tcc.W
			h1 := tcc.H

			if physics.Collide(x1, y1, w1, h1, x2, y2, w2, h2) {
				laserEnt.SetActive(false)
				laserEnt.GetEngine().Kill(v)
			}
		}
	}

}

func (lms *LaserMoveSystem) Draw(screen *ebiten.Image) {}

func (lms *LaserMoveSystem) Active() bool {
	return lms.active
}

func (lms *LaserMoveSystem) SetActive(active bool) {
	lms.active = active
}

func (lms *LaserMoveSystem) AddEntityIfRequired(e types.IEntity) {
	if _, ok := lms.targets[e.GetID()]; ok {
		return
	}

	for _, c := range lms.filter.Requires() {
		if _, ok := e.GetComponents()[c]; !ok {
			return
		}
	}

	logger.Debug("System %T added entity %d ", lms, e.GetID())

	lms.targets[e.GetID()] = e
}

func (lms *LaserMoveSystem) RemoveEntityIfRequired(e types.IEntity) {
	for _, c := range lms.filter.Requires() {
		if !e.HasComponent(c) {
			logger.Debug("System %T removed entity %d ", lms, e.GetID())

			delete(lms.targets, e.GetID())

			return
		}
	}
}

func (s *LaserMoveSystem) RemoveEntity(e types.IEntity) {
	logger.Debug("System %T removed entity %d ", s, e.GetID())

	delete(s.targets, e.GetID())
}
