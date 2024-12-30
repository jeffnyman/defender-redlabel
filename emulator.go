package main

import (
	"errors"
	"fmt"
	_ "image/png"

	"github.com/jeffnyman/defender-redlabel/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Emulator struct {
	engine *game.Engine
}

func NewApp() *Emulator {
	emulator := &Emulator{
		engine: game.NewEngine(),
	}
	emulator.engine.Init()
	return emulator
}

func (emulator *Emulator) Update() error {

	if status := emulator.engine.Update(); status != game.OK {
		return errors.New(fmt.Sprintf("%d", status))
	}
	return nil
}

func (emulator *Emulator) Draw(screen *ebiten.Image) {
	emulator.engine.Draw(screen)
}

func (g *Emulator) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}
