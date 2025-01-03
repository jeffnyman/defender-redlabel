package player

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/event"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerPlay struct {
	Name types.StateType
}

func NewPlayerPlay() *PlayerPlay {
	return &PlayerPlay{
		Name: types.PlayerPlay,
	}
}

func (s *PlayerPlay) GetName() types.StateType {
	return s.Name
}

func (s *PlayerPlay) Enter(ai *components.AI, e types.IEntity) {
	sc := e.GetComponent(types.Ship).(*components.Ship)
	sc.ScreenOffset = defs.ScreenWidth * 0.1
	ev := event.NewStart(e)
	event.NotifyEvent(ev)
}

func (s *PlayerPlay) Update(ai *components.AI, e types.IEntity) {
	pc := e.GetComponent(types.Pos).(*components.Pos)
	sc := e.GetComponent(types.Ship).(*components.Ship)
	dc := e.GetComponent(types.Draw).(*components.Draw)

	if sc.Direction == 1 && sc.ScreenOffset > defs.ScreenWidth*0.1 {
		sc.ScreenOffset -= 30
	}

	if sc.Direction == -1 && sc.ScreenOffset < defs.ScreenWidth*0.9 {
		sc.ScreenOffset += 30
	}

	camx := (pc.X - sc.ScreenOffset)

	if camx < 0 {
		camx += defs.WorldWidth
	}

	defs.SetCameraX(camx)

	fle := e.GetEngine().GetEntity(e.Child())
	flpc := fle.GetComponent(types.Pos).(*components.Pos)
	fldc := fle.GetComponent(types.Draw).(*components.Draw)

	flpc.X = pc.X - 40

	if sc.Direction < 0 {
		flpc.X = pc.X + 5
	}

	flpc.Y = pc.Y + 10

	fdc := fle.GetComponent(types.Draw).(*components.Draw)

	if ebiten.IsKeyPressed(defs.KeyMap[types.Reverse]) {
		if !sc.ReversePressed {
			sc.Direction = -sc.Direction
			dc.FlipX = !dc.FlipX
			fdc.FlipX = !fdc.FlipX
			sc.ReversePressed = true
			pc.DX /= 2
		}
	} else {
		sc.ReversePressed = false
	}

	if ebiten.IsKeyPressed(defs.KeyMap[types.Thrust]) {
		if !sc.ThrustPressed {
			sc.ThrustPressed = true
			ev := event.NewPlayerThrust(e)
			event.NotifyEvent(ev)
			fldc.Hide = false
		}
		pc.DX += sc.Direction * 2
	} else {
		if sc.ThrustPressed {
			ev := event.NewPlayerStopThrust(e)
			event.NotifyEvent(ev)
		}

		fldc.Hide = true
		sc.ThrustPressed = false
		pc.DX /= 1.05
	}

	if pc.DX > defs.PlayerSpeedX {
		pc.DX = defs.PlayerSpeedX
	}

	if pc.DX < -defs.PlayerSpeedX {
		pc.DX = -defs.PlayerSpeedX
	}

	if ebiten.IsKeyPressed(defs.KeyMap[types.Up]) && ebiten.IsKeyPressed(defs.KeyMap[types.Down]) {
		pc.DY = 0
	} else if ebiten.IsKeyPressed(defs.KeyMap[types.Up]) {
		if pc.DY > -defs.PlayerSpeedY {
			pc.DY -= 2
		}
	} else if ebiten.IsKeyPressed(defs.KeyMap[types.Down]) {
		if pc.DY < defs.PlayerSpeedY {
			pc.DY += 2
		}
	} else {
		pc.DY = 0
	}

	if ebiten.IsKeyPressed(defs.KeyMap[types.Fire]) {
		if !sc.FirePressed {
			sc.FirePressed = true
			ev := event.NewPlayerFire(e)
			event.NotifyEvent(ev)
		}
	} else {
		sc.FirePressed = false
	}

	if ebiten.IsKeyPressed(defs.KeyMap[types.SmartBomb]) {
		if defs.SmartBombs == 0 {
			return
		}
		if !sc.SmartBombPressed {
			sc.SmartBombPressed = true
			ev := event.NewSmartBomb(pc)
			event.NotifyEvent(ev)
			event.NotifyEventDelay(ev, 5)
			event.NotifyEventDelay(ev, 10)
			defs.SmartBombs--
		}
	} else {
		sc.SmartBombPressed = false
	}

	if ebiten.IsKeyPressed(defs.KeyMap[types.HyperSpace]) {
		if !sc.HyperSpacePressed {
			sc.HyperSpacePressed = true
			ev := event.NewPlayerDie(e)
			event.NotifyEvent(ev)
		}
	} else {
		sc.HyperSpacePressed = false
	}
}
