package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/types"
)

type GameIntro struct {
	Name types.StateType
}

func NewGameIntro() *GameIntro {
	return &GameIntro{
		Name: types.GameIntro,
	}
}

func (s *GameIntro) GetName() types.StateType {
	return s.Name
}

func (s *GameIntro) Enter(ai *cmp.AI, e types.IEntity) {
}

func (s *GameIntro) Update(ai *cmp.AI, e types.IEntity) {
	ai.NextState = types.GameStart
}
