package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/types"
)

type GameStart struct {
	Name types.StateType
}

func NewGameStart() *GameStart {
	return &GameStart{
		Name: types.GameStart,
	}
}

func (s *GameStart) GetName() types.StateType {
	return s.Name
}

func (s *GameStart) Enter(ai *cmp.AI, e types.IEntity) {
	e.GetEngine().LevelStart()
}

func (s *GameStart) Update(ai *cmp.AI, e types.IEntity) {
	ai.NextState = types.GamePlay
}
