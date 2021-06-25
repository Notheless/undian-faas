package undian

import (
	"context"
)

func (s *service) CleanPemenang(ctx context.Context) error {
	sql := `UPDATE pemenang SET deleted = 1`
	err := s.db.QueryRowContext(ctx, sql).Err()
	if err != nil {
		return err
	}
	return nil
}
