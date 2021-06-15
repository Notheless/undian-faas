package undian

import (
	"context"
	"database/sql"
)

type (
	// Service interface
	Service interface {
		GetUndian(ctx context.Context) error
	}

	service struct {
		db *sql.DB
	}
)

// NewService func
func NewService(db *sql.DB) Service {
	return &service{db: db}
}
