package drivechecker

// DriveState struct to represent a store state
type DriveState struct {
	IsActive bool
	Dispo    string
}

type state map[string]*DriveState

var currentState = make(state)

// GetDriveState get the state of a drive
func GetDriveState(storeID string) *DriveState {
	return currentState[storeID]
}

// SetDriveState create a new state
func SetDriveState(storeID string, state *DriveState) {
	currentState[storeID] = state
}
