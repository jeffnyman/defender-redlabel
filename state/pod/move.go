package pod

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/types"
)

type PodMove struct {
	Name types.StateType
}

func NewPodMove() *PodMove {
	return &PodMove{
		Name: types.PodMove,
	}
}

func (s *PodMove) GetName() types.StateType {
	return s.Name
}

func (s *PodMove) Enter(ai *components.AI, e types.IEntity) {
}

func (s *PodMove) Update(ai *components.AI, e types.IEntity) {
}
