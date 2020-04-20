package drivechecker

// StoreConfig struct reprsent a Store config
type StoreConfig struct {
	StoreID string
	State   *DriveState
}

// NewStoreConfig Create a new Store config with storeID
func NewStoreConfig(storeID string) StoreConfig {
	return StoreConfig{
		StoreID: storeID,
		State:   &DriveState{},
	}
}
