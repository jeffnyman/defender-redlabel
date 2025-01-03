package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/components"
	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type GameLevelEnd struct {
	Name types.StateType
}

func NewGameLevelEnd() *GameLevelEnd {
	return &GameLevelEnd{
		Name: types.GameLevelEnd,
	}
}

func (s *GameLevelEnd) GetName() types.StateType {
	return s.Name
}

func (s *GameLevelEnd) Enter(ai *components.AI, e types.IEntity) {
	e.GetEngine().LevelEnd()
	ai.Scratch = 0
}

func (s *GameLevelEnd) Update(ai *components.AI, e types.IEntity) {
	ai.Scratch++

	if ai.Scratch > 4*30 {
		defs.NextLevel()
		ai.NextState = types.GameStart
	}
}
