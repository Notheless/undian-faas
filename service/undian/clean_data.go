package undian

import (
	"context"
)

func (s *service) CleanPemenang(ctx context.Context) error {
	sql := `UPDATE pemenang SET deleted = 1`

	_, err := s.db.QueryContext(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}
