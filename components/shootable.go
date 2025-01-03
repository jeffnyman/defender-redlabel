package components

import "github.com/jeffnyman/defender-redlabel/types"

type Shootable struct {
	componentType types.CmpType
}

func NewShootable() *Shootable {
	return &Shootable{
		componentType: types.Shootable,
	}
}

func (pos *Shootable) Type() types.CmpType {
	return pos.componentType
}
