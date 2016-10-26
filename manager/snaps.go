package manager

import (
	"time"

	"github.com/ubuntu/snaps-manager/state"
)

type snapState struct {
	state.State
	WaitTime time.Duration
	Action   state.Transition
}
