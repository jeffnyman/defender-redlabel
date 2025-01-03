package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/components"
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

func (s *GameIntro) Enter(ai *components.AI, e types.IEntity) {
}

func (s *GameIntro) Update(ai *components.AI, e types.IEntity) {
	ai.NextState = types.GameStart
}
