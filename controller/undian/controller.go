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

// NewController func
func Warp(srv undian.Service) *Controller {
	return &Controller{srv: srv}
}

// Route func
func (c *Controller) Route(r *mux.Router) {
	s := r.PathPrefix("/undian").Subrouter()
	s.HandleFunc("/{zona}/{kategori}", c.GeneratePemenang).Methods("POST")
	s.HandleFunc("/{zona}/{kategori}", c.LihatPemenang).Methods("GET")
	s.HandleFunc("/history", c.LihatSemuaPemenang).Methods("GET")
	s.HandleFunc("/kategori", c.LihatSemuaKategori).Methods("GET")
	s.HandleFunc("/zona", c.LihatSemuaZona).Methods("GET")
}