package components

import "github.com/jeffnyman/defender-redlabel/types"

type LaserMove struct {
	componentType types.CmpType
	Length        float64
}

func NewLaserMove() *LaserMove {
	return &LaserMove{
		componentType: types.LaserMove,
	}
}

func (pos *LaserMove) Type() types.CmpType {
	return pos.componentType
}
