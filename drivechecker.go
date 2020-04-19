package drivechecker

import (
	"log"
	"time"

	"github.com/nlevee/go-drive-checker/pkg/drivestate"
)

type DriveChecker interface {
	GetStoreByPostalCode(postalCode string) ([]DriveStore, error)
	LoadDriveState(config DriveConfig) (hasChanged bool, err error)
}

type DriveConfig struct {
	DriveID string
	State   *drivestate.DriveState
}

type DriveStore struct {
	DriveID string
	Name    string
}

// NewConfig Create a new Drive config with driveId
func NewConfig(driveID string) DriveConfig {
	state := &drivestate.DriveState{
		IsActive: false,
		Dispo:    "",
	}
	return DriveConfig{
		DriveID: driveID,
		State:   state,
	}
}

// GetStoreIDByPostalCode fetch storeIDs by postal code
func GetStoreIDByPostalCode(drive DriveChecker, postalCode string) ([]string, error) {
	storeIds := []string{}

	stores, err := drive.GetStoreByPostalCode(postalCode)
	if err != nil {
		return storeIds, err
	}

	for _, v := range stores {
		storeIds = append(storeIds, v.DriveID)
	}

	return storeIds, nil
}

// LoadIntervalDriveState fetch each tick the drive state config
func LoadIntervalDriveState(drive DriveChecker, config DriveConfig, tick *time.Ticker, done chan bool) {
	log.Printf("Démarrage du check de créneau %v", config.DriveID)

	// premier appel sans attendre le premier tick
	if _, err := drive.LoadDriveState(config); err != nil {
		log.Print(err)
	}

	for {
		select {
		case <-tick.C:
			// a chaque tick du timer on lance une recherche de state
			if _, err := drive.LoadDriveState(config); err != nil {
				log.Print(err)
			}
		case <-done:
			log.Printf("Ticker stopped")
			tick.Stop()
			return
		}
	}
}

// GetDriveState get the state of a drive
func GetDriveState(driveID string) *drivestate.DriveState {
	return drivestate.GetDriveState(driveID)
}

// NewDriveHandler add a new drive handler
func NewDriveHandler(drive DriveChecker, driveID string) {
	config := NewConfig(driveID)
	drivestate.NewDriveState(driveID, config.State)

	tick := time.NewTicker(2 * time.Minute)
	done := make(chan bool)

	LoadIntervalDriveState(drive, config, tick, done)
}
