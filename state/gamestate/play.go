package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/components"

	"github.com/jeffnyman/defender-redlabel/defs"
	"github.com/jeffnyman/defender-redlabel/types"
)

type GamePlay struct {
	Name types.StateType
}

func NewGamePlay() *GamePlay {
	return &GamePlay{
		Name: types.GamePlay,
	}
}

func (s *GamePlay) GetName() types.StateType {
	return s.Name
}

func (s *GamePlay) Enter(ai *components.AI, e types.IEntity) {
}

func (s *GamePlay) Update(ai *components.AI, e types.IEntity) {

	if defs.LandersKilled == defs.CurrentLevel().LanderCount {
		ai.NextState = types.GameLevelEnd
	}
	if defs.PlayerLives == 0 {
		ai.NextState = types.GameOver
	}

}
