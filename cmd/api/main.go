package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	drivechecker "github.com/nlevee/go-drive-checker"
)

// Env
type Env struct {
	Retail *drivechecker.Retail
}

// StartServer demarrage du server
func StartServer(e Env, host string, port string) {
	r := mux.NewRouter()
	e.addStoreRoutes(r)
	e.addScrapperRoutes(r)
	log.Fatal(http.ListenAndServe(host+":"+port, r))
}
