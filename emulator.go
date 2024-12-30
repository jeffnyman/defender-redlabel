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
	app := &Emulator{
		engine: game.NewEngine(),
	}
	app.engine.Init()
	return app
}

func (app *Emulator) Update() error {

	if status := app.engine.Update(); status != game.OK {
		return errors.New(fmt.Sprintf("%d", status))
	}
	return nil
}

func (app *Emulator) Draw(screen *ebiten.Image) {
	app.engine.Draw(screen)
}

func (g *Emulator) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320 * 5, 240 * 5
}
