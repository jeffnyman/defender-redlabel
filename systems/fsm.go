package systems

import (
	"fmt"

	"github.com/jeffnyman/defender-redlabel/cmp"
	"github.com/jeffnyman/defender-redlabel/logger"
	"github.com/jeffnyman/defender-redlabel/types"
)

var fsmList []FSM
var fsmCount int

func init() {
	fsmList = []FSM{}
	fsmCount = 0
}

type FSM struct {
	stategraph *StateGraph
}

func NewFSM(s *StateGraph) int {
	fsmList = append(fsmList, FSM{
		stategraph: s,
	})

	rv := fsmCount
	fsmCount++

	return rv
}

func GetFSM(id int) FSM {
	return fsmList[id]
}

func (f FSM) Update(ai *cmp.AI, e types.IEntity) {
	if ai.NextState != ai.State {
		next_state, err := f.stategraph.State(ai.NextState)

		if err != nil {
			panic(fmt.Sprintf("No state defined in FSM for %s", ai.NextState.String()))
		}

		next_state.Enter(ai, e)
		logger.Debug("Entity %d state change %s -> %s", e.GetID(), ai.State.String(), ai.NextState.String())

		ai.State = ai.NextState
	}

	curr_state, err := f.stategraph.State(ai.State)

	if err != nil {
		panic(fmt.Sprintf("no current state %s in FSM ", ai.State.String()))
	}

	curr_state.Update(ai, e)
}
