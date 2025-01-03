package components

import (
	"github.com/jeffnyman/defender-redlabel/graphics"
	"github.com/jeffnyman/defender-redlabel/types"

	"github.com/hajimehoshi/ebiten/v2"
)

type Draw struct {
	componentType types.CmpType
	Image         *ebiten.Image
	Opts          *ebiten.DrawImageOptions
	Color         types.ColorF
	Scale         float64
	SpriteMap     graphics.GFXFrame
	Counter       int
	Frame         int
	Disperse      float64
	Cycle         bool
	CycleIndex    float64
	Bomber        bool
	FlipX         bool
	Hide          bool
}

func NewDraw(image *ebiten.Image, smap graphics.GFXFrame, color types.ColorF) *Draw {
	return &Draw{
		Image:         image,
		Opts:          &ebiten.DrawImageOptions{},
		Color:         color,
		componentType: types.Draw,
		Scale:         1,
		SpriteMap:     smap,
		Counter:       0,
		Frame:         0,
		Disperse:      0,
		CycleIndex:    0,
		Bomber:        false,
		FlipX:         false,
		Hide:          false,
	}
}

func (d *Draw) Type() types.CmpType {
	return d.componentType
}
