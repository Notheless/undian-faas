package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	undianctrl "main/controller/undian"
	undiansrv "main/service/undian"
	util "main/util"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	godotenv.Load()
	db, err := util.NewDBClient()
	if err != nil {
		log.Fatal(err)
	}

	// create mux router
	r := mux.NewRouter()

	// create master data mitra controller
	undianctrl.NewController(undiansrv.NewService(db)).Route(r)
	port := ":" + os.Getenv("APP_PORT")
	fmt.Println(port)
	// run http server
	if err := http.ListenAndServe(port, handlers.RecoveryHandler()(r)); err != nil {
		panic(err)
	}
}
