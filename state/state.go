package state

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/types"
)

type IState interface {
	GetName() types.StateType
	Enter(ai *components.AI, e types.IEntity)
	Update(ai *components.AI, e types.IEntity)
}
