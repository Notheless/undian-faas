package undian

import (
	"context"
	"database/sql"
)

type (
	// Service interface
	Service interface {
		GeneratePemenang(ctx context.Context, zona string, kategori string) ([]PemenangResult, error)
		LihatSemuaKategori(ctx context.Context) ([]KategoriResult, error)
		LihatSemuaZona(ctx context.Context) ([]string, error)
		LihatPemenang(ctx context.Context, zona string, kategori string) ([]PemenangResult, error)
		LihatPemenangZonasi(ctx context.Context, zona string) ([]PemenangZonaResult, error)
		LihatSemuaPemenang(ctx context.Context) ([]PemenangSemuaResult, error)

		LihatPemenangQuery(ctx context.Context, zonaQ []string, kategoriQ []string) ([]PemenangSemuaResult, error)
		ExportExcel(ctx context.Context) (string, error)
		GeneratePemenangQuery(ctx context.Context, zonaQ []string, kategoriQ []string) ([]PemenangSemuaResult, error)
	}

	service struct {
		db *sql.DB
	}

	KategoriResult struct {
		Nama        string `json:"nama"`
		Description string `json:"description"`
		Jumlah      int    `json:"jumlah"`
	}

	PemenangResult struct {
		Tiket    string `json:"tiket"`
		NamaToko string `json:"nama_toko"`
	}

	PemenangZonaResult struct {
		Kategori            string           `json:"kategori"`
		KategoriDescription string           `json:"kategori_desc"`
		Pemenang            []PemenangResult `json:"pemenang"`
	}

	PemenangSemuaResult struct {
		Zona         string               `json:"zona"`
		ZonaPemenang []PemenangZonaResult `json:"zona_pemenang"`
	}

	temp struct {
		Zona string
	}
)

// NewService func
func NewService(db *sql.DB) Service {
	return &service{db: db}
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
