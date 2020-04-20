package drivechecker

import (
	"log"
	"time"
)

// DriveScrapper interface that defines func to get drive state
type DriveScrapper interface {
	LoadDriveState(config StoreConfig) (hasChanged bool, err error)
}

// Store struct to represent a store data
type Store struct {
	ID   string
	Name string
	DriveScrapper
}

// LoadIntervalDriveState fetch each tick the drive state config
func (s Store) LoadIntervalDriveState(config StoreConfig, tick *time.Ticker, done chan bool) {
	log.Printf("Démarrage du check de créneau %v", config.StoreID)

	// premier appel sans attendre le premier tick
	if _, err := s.LoadDriveState(config); err != nil {
		log.Print(err)
	}

	for {
		select {
		case <-tick.C:
			// a chaque tick du timer on lance une recherche de state
			if _, err := s.LoadDriveState(config); err != nil {
				log.Print(err)
			}
		case <-done:
			log.Printf("Ticker stopped")
			tick.Stop()
			return
		}
	}
}

// NewDriveHandler add a new drive handler
func (s Store) NewDriveHandler() {
	config := NewStoreConfig(s.ID)
	SetDriveState(s.ID, config.State)

	tick := time.NewTicker(2 * time.Minute)
	done := make(chan bool)

	s.LoadIntervalDriveState(config, tick, done)
}
