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
	s := r.PathPrefix("/undian").Subrouter()
	s.HandleFunc("", c.LihatPemenangQuery).Methods("GET")
	s.HandleFunc("", c.GeneratePemenangQuery).Methods("POST")
	s.HandleFunc("/kategori", c.LihatSemuaKategori).Methods("GET")
	s.HandleFunc("/zona", c.LihatSemuaZona).Methods("GET")
	s.HandleFunc("/history", c.LihatSemuaPemenang).Methods("GET")
	s.HandleFunc("/export", c.ExportExcel).Methods("GET")
	s.HandleFunc("/{zona}", c.LihatPemenangZonasi).Methods("GET")
	s.HandleFunc("/{zona}/{kategori}", c.LihatPemenang).Methods("GET")
	s.HandleFunc("/{zona}/{kategori}", c.GeneratePemenang).Methods("POST")
}
