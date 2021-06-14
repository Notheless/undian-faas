package undian

import (
	"main/service/undian"

	"github.com/gorilla/mux"
)

// Controller struct
type Controller struct {
	srv undian.Service
}

// NewController func
func NewController(srv undian.Service) *Controller {
	return &Controller{srv: srv}
}

// Route func
func (c *Controller) Route(r *mux.Router) {
	s := r.PathPrefix("/master-data/mitra").Subrouter()
	s.HandleFunc("", c.GetUndian).Methods("GET")
}
