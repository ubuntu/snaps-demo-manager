package manager

import (
	"sync"
	"time"

	"github.com/ubuntu/snaps-demo-manager/logger"
	"github.com/ubuntu/snaps-demo-manager/state"
)

type snapState struct {
	state.State
	WaitTime time.Duration
	Action   state.Transition
}

// Snap on the system.
type Snap struct {
	Name         string
	currentState snapState
	Instructions chan []state.SnapInstruction
	Stop         chan bool
}

// NewSnap initializes a new snap object.
func NewSnap(name string) *Snap {
	logger.Info("Prepare new snap %s", name)
	s := &Snap{Name: name}
	s.Instructions = make(chan []state.SnapInstruction)
	s.Stop = make(chan bool)
	return s
}

// Track and manage snap lifecycle.
func (s *Snap) Track(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer logger.Info("Stop tracking %s", s.Name)

		logger.Info("Tracking %s", s.Name)
		for {
			select {
			case <-s.Instructions:
				logger.Info("New set of instructions received for %s", s.Name)

			case <-s.Stop:
				return
			}
		}
	}()
}
