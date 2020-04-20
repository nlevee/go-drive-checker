package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	drivechecker "github.com/nlevee/go-drive-checker"
)

// AddScrapperRoutes populate router
func (e Env) addScrapperRoutes(r *mux.Router) {
	// ajoute un scrapper sur le store : storeid
	r.HandleFunc("/scrappers/{storeid}", e.addScrapper).Methods(http.MethodPut)
	// récupère l'état du scrapper sur le store : storeid
	r.HandleFunc("/scrappers/{storeid}", e.getScrapperState).Methods(http.MethodGet)
}

// GetScrapperState récuperation d'un état de scrapper
func (e Env) getScrapperState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	log.Printf("storeId: %v\n", vars["storeid"])

	json.NewEncoder(w).Encode(drivechecker.GetDriveState(vars["storeid"]))
}

// AddScrapper ajoute un scrapper
func (e Env) addScrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	vars := mux.Vars(r)
	log.Printf("storeId: %v\n", vars["storeid"])

	store, _ := e.Retail.GetStoreById(vars["storeid"])
	go store.NewDriveHandler()
}
