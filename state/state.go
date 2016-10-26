package state

// State that can be a snap in
type State struct {
	Channel  string
	DevMode  bool
	Revision string
	Version  string
}

// Transition from one state to another one
type Transition struct {
	Action Action
	Next   *State
}

// Action associated to a Transition
type Action int

const (
	// Install or update a snap
	Install = iota
	// Revert a snap
	Revert
	// Remove a snap
	Remove
)
