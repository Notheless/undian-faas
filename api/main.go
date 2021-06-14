package main

import (
	"log"

	"net/http"

	undianctrl "main/controller/undian"
	undiansrv "main/service/undian"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db, err := NewDBClient()
	if err != nil {
		log.Fatal(err)
	}

	// create mux router
	r := mux.NewRouter()

	// create master data mitra controller
	undianctrl.NewController(undiansrv.NewService(db)).Route(r)

	// run http server
	if err := http.ListenAndServe(":8080", handlers.RecoveryHandler()(r)); err != nil {
		panic(err)
	}
}
