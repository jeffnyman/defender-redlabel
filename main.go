package main

import (
	_ "image/png"
	"math/rand"
	"time"

	"github.com/jeffnyman/defender-redlabel/defs"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	ebiten.SetWindowSize(320*5, 240*5)
	ebiten.SetWindowTitle("Defender (RedLabel)")
	ebiten.SetFullscreen(true)
	ebiten.SetMaxTPS(defs.MaxTPS)

	emulator := NewEmulator()
	if err := ebiten.RunGame(emulator); err != nil {
	}
}
