package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/types"
)

type GameOver struct {
	Name types.StateType
}

func NewGameOver() *GameOver {
	return &GameOver{
		Name: types.GameOver,
	}
}

func (s *GameOver) GetName() types.StateType {
	return s.Name
}

func (s *GameOver) Enter(ai *components.AI, e types.IEntity) {
	ai.Scratch = 0
	e.GetEngine().GameOver()
}

func (s *GameOver) Update(ai *components.AI, e types.IEntity) {
	ai.Scratch++

	if ai.Scratch > 4*30 {
		e.GetEngine().Terminate()
	}
}
