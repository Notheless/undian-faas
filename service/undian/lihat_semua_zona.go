package undian

import (
	"context"
)

func (s *service) LihatSemuaZona(ctx context.Context) ([]string, error) {
	var res []string

	sql := `SELECT id FROM undian.zona`
	rs, err := s.db.QueryContext(ctx, sql)
	if err != nil {
		return res, err
	}
	for rs.Next() {
		var str string
		rs.Scan(&str)
		res = append(res, str)
	}
	return res, nil
}
