package state

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/types"
)

type IState interface {
	GetName() types.StateType
	Enter(ai *cmp.AI, e types.IEntity)
	Update(ai *cmp.AI, e types.IEntity)
}
