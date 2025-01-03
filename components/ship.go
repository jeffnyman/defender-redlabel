package components

import (
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type Ship struct {
	componentType     types.CmpType
	Direction         float64
	ScreenOffset      float64
	ReversePressed    bool
	FirePressed       bool
	SmartBombPressed  bool
	HyperSpacePressed bool
	ThrustPressed     bool
	W, H              float64
}

func NewShip(w, h int) *Ship {
	return &Ship{
		ScreenOffset:  defs.ScreenWidth * 0.2,
		Direction:     1,
		componentType: types.Ship,
		W:             float64(w),
		H:             float64(h),
	}
}

func (pos *Ship) Type() types.CmpType {
	return pos.componentType
}
