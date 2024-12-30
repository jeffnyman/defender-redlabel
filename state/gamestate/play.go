package gamestate

import (
	"github.com/jeffnyman/defender-redlabel/cmp"

	"github.com/jeffnyman/defender-redlabel/gl"
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

func (s *GamePlay) Enter(ai *cmp.AI, e types.IEntity) {
}

func (s *GamePlay) Update(ai *cmp.AI, e types.IEntity) {

	if gl.LandersKilled == gl.CurrentLevel().LanderCount {
		ai.NextState = types.GameLevelEnd
	}
	if gl.PlayerLives == 0 {
		ai.NextState = types.GameOver
	}

}
