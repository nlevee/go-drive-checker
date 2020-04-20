package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddStoreRoutes populate router
func (e Env) addStoreRoutes(r *mux.Router) {
	// récuperation des stores
	r.HandleFunc("/stores", e.getStores).Methods(http.MethodGet)
}

// GetStores récupere la liste des stores filtrés
func (e Env) getStores(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cp := params["postalCode"]
	if len(cp) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	storeIDs, _ := e.Retail.GetStoreByPostalCode(string(cp[0]))

	json.NewEncoder(w).Encode(storeIDs)
}
